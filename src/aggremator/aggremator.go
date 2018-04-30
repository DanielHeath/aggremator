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
TODO: Extract all these feeds to a config file or something
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
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

var mailer *gomail.Mailer

func init() {
	if os.Getenv("AGGREMATOR_PW") == "" {
		panic("Must set AGGREMATOR_PW")
	}

	mailer = gomail.NewMailer(
		"mail.gandi.net",
		"rss@nerdy.party",
		os.Getenv("AGGREMATOR_PW"),
		465,
	)
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
				msg.SetHeader("From", "rss@nerdy.party")
				msg.SetHeader("To", "rss@nerdy.party")
				msg.SetHeader("Subject", "Error")
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
			if _, alreadyFound := pastEntries[maildir.CleanId(item.Link+item.ID)]; item.Link != "" && !alreadyFound {
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss@nerdy.party")
				msg.SetHeader("To", "rss@nerdy.party")
				err := feed.Serialize(*item, msg)

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
	log.Printf("Sending %+v\n", msg)
	die(mailer.Send(msg))
	log.Printf("Sent\n")
}

func die(err error, context ...interface{}) {
	if err != nil {
		fmt.Println(context)
		panic(err)
	}
}
