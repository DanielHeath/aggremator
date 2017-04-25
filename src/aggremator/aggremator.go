package main

/*
TODO:
  * PA -> title has escaped entities
	* write 'go generate' based curl-to-sample thingy instead of manually updating
	* I'd prefer to extract the feeds to a config file so this could be re-usable
	* amazingsuperpowers had a youtube link in their feed, came thru as a busted thingy (using an iframe)
	* QC responded to a comic link with a 404 (not up yet?) but no error was logged.
  * background test: hit a standard item from the feed, apply selector, check you get the right output
   * that way you can find out when the layout breaks your selector...
*/

/*
TODO:
 * Uplift to the cloud
 * gopherjs -> lambda?
 * https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words
 *
 * s3://aggremator/<site>/feedurl -> 10-day retention w/
 *
 * structure:
 * -> feed GET (curl) - input=url, output=xml
 * -> change detection (new entries) - inputs=xml+pastentries, outputs=[]url|[]message
 * -> entry processing (fetch/transform) - inputs=url, output=message
 * -> delivery (email) - inputs=message, output=smtp|offlineimapfile|other
 *
 * pastentries dataset size: 1kb/url, 30 urls/month * 10 feeds -> ~4mb/year. Will fit into ram on a tiny node, probably forever.
 * Can prune every 10y to avoid having to manage it better.
*/

// I already know how to manage (a) server.
// though it's nice to not have to reinstall everything when re-imaging, that's a ~45 minute saving in a theoretical future
// vs several hours now.
// put it on a spare email account (nerdy.party?) and add it to clients?

import (
	"flag"
	"fmt"
	"os/user"
	"reflect"

	"aggremator/feeds"
	"aggremator/feeds/alicegrove"
	"aggremator/feeds/amazingsuperpowers"
	"aggremator/feeds/codelesscode"
	"aggremator/feeds/dilbert"
	"aggremator/feeds/flashboardwars"
	"aggremator/feeds/fowllanguage"
	"aggremator/feeds/orderofthestick"
	"aggremator/feeds/pbfcomics"
	"aggremator/feeds/pennyarcade"
	"aggremator/feeds/questionablecontent"
	"aggremator/feeds/sarahs_doodles"
	"aggremator/feeds/smbc"
	"aggremator/feeds/webcomicname"
	"aggremator/feeds/whatif"
	"aggremator/feeds/xkcd"
	"aggremator/maildir"
	"aggremator/pastentries"

	"github.com/SlyMarbo/rss"
	multierror "github.com/hashicorp/go-multierror"
	gomail "gopkg.in/gomail.v1"
)

var pastEntriesPath string
var homedir string
var debug bool
var maildirPath string

func init() {
	usr, err := user.Current()
	die(err)
	homedir = usr.HomeDir
	flag.StringVar(
		&pastEntriesPath,
		"pastEntriesFile",
		homedir+"/.aggremator/pastentries",
		"File to store which feed items have already been sync-ed",
	)
	flag.BoolVar(
		&debug,
		"debug",
		false,
		"debug mode",
	)
}

var allFeeds = []feeds.Feed{
	xkcd.Feed,
	pennyarcade.Feed,
	dilbert.Feed,
	alicegrove.Feed,
	amazingsuperpowers.Feed,
	sarahs_doodles.Feed,
	webcomicname.Feed,
	fowllanguage.Feed,
	flashboardwars.Feed,
	smbc.Feed,
	questionablecontent.Feed,
	whatif.Feed,
	orderofthestick.Feed,
	codelesscode.Feed,
	pbfcomics.Feed,
}
var currentFeeds = allFeeds

func main() {
	var err error
	flag.Parse()

	pastEntries := pastentries.PastEntries{}
	pastEntriesFile := pastentries.File(pastEntriesPath)
	if !debug {
		pastEntries, err = pastEntriesFile.Read()
		die(err)
	}

	maildirPath = homedir + "/.mail/fastmail"
	if debug {
		maildirPath = homedir + "/.mail/testmail"
	}

	var feedErrors *multierror.Error

	for _, feed := range currentFeeds {
		fail := func(err error) bool {
			if err != nil {
				errType := reflect.TypeOf(err)

				// Sometimes hosts are down, causing the 'net' package to return errors.
				// We'll try again later; no need to log this.
				if errType.PkgPath() == "net" {
					return true
				}

				feedErrors = multierror.Append(feedErrors, err)
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss.errors@example.org")
				msg.SetHeader("To", "rss.errors@example.org")
				msg.SetBody("text/plain", err.Error()+"\n"+feed.Url()+"\n"+errType.PkgPath()+"\n"+errType.String())
				send(
					msg,

					maildirPath+fmt.Sprintf(
						"/INBOX.Errors.%s/new/%d",
						feed.Category(),
						&err,
					),
				)

				return true
			}
			return false
		}

		var doc *rss.Feed
		if debug {
			doc, err = rss.Parse([]byte(feed.Sample()))
		} else {
			doc, err = rss.Fetch(feed.Url())
		}
		if fail(err) {
			continue
		}
		for _, item := range doc.Items {

			// TODO: One goroutine per feed? one process per feed?
			// Have we seen this feed entry before?
			// TODO: pastEntries should be a smarter type, incorporating CleanId, not just a map; also, one-per-feed?
			if _, ok := pastEntries[maildir.CleanId(item.Link+item.ID)]; !ok {
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss@example.org")
				msg.SetHeader("To", "rss@example.org")
				err := feed.Serialize(*item, msg)
				// die(err, item, "\n", item.Content)
				if fail(err) {
					continue
				}

				send(
					msg,
					maildirPath+
						"/INBOX.Feeds."+
						feed.Category()+
						"/new/"+
						maildir.CleanId(item.ID),
				)

				if !debug {
					pastEntries[maildir.CleanId(item.Link+item.ID)] = true
					die(pastEntries.Write(pastEntriesFile)) // Update past entries after each message.
				}

			}
		}
	}
	if feedErrors != nil {
		fmt.Println(feedErrors.Errors)
		fmt.Println(feedErrors.GoString())
		fmt.Println(feedErrors.WrappedErrors())
		die(feedErrors.ErrorOrNil())
	}
}

func send(msg *gomail.Message, sendPath string) {
	// TODO: SMTP directly (CLI options for sender/recipient/credentials?)
	if debug {
		die(gomail.NewCustomMailer("127.0.0.1:1025", nil).Send(msg))
	} else {
		sender := maildir.Mailer(sendPath)

		m := gomail.NewMailer("localhost", "dummy", "dummy", 9002, gomail.SetSendMail(sender))
		die(m.Send(msg))
	}
}

func die(err error, context ...interface{}) {
	if err != nil {
		fmt.Println(context)
		panic(err)
	}
}
