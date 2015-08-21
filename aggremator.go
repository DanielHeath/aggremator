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

import (
	"flag"
	"fmt"
	"net"
	"os/user"
	"strings"

	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds"
	"github.com/danielheath/aggremator/feeds/alicegrove"
	"github.com/danielheath/aggremator/feeds/amazingsuperpowers"
	"github.com/danielheath/aggremator/feeds/codelesscode"
	"github.com/danielheath/aggremator/feeds/dilbert"
	"github.com/danielheath/aggremator/feeds/orderofthestick"
	"github.com/danielheath/aggremator/feeds/pbfcomics"
	"github.com/danielheath/aggremator/feeds/pennyarcade"
	"github.com/danielheath/aggremator/feeds/questionablecontent"
	"github.com/danielheath/aggremator/feeds/smbc"
	"github.com/danielheath/aggremator/feeds/whatif"
	"github.com/danielheath/aggremator/feeds/xkcd"
	"github.com/danielheath/aggremator/maildir"
	"github.com/danielheath/aggremator/pastentries"
	"github.com/go-gomail/gomail"
	"github.com/hashicorp/go-multierror"
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
				if _, ok := err.(net.Error); ok {
					return true
				}
				if _, ok := err.(net.DNSError); ok {
					return true
				}
				if _, ok := err.(net.AddrError); ok {
					return true
				}
				if _, ok := err.(net.DNSConfigError); ok {
					return true
				}
				if _, ok := err.(net.InvalidAddrError); ok {
					return true
				}
				if _, ok := err.(net.OpError); ok {
					return true
				}
				if _, ok := err.(net.ParseError); ok {
					return true
				}

				feedErrors = multierror.Append(feedErrors, err)
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss.errors@example.org")
				msg.SetHeader("To", "rss.errors@example.org")
				msg.SetBody("text/plain", err.Error()+"\n"+feed.Url())
				send(
					msg,

					maildirPath+fmt.Sprintf(
						"/INBOX.Feeds.%s/new/%d",
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

			// TODO: One goroutine per feed
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
	die(feedErrors.ErrorOrNil())
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
