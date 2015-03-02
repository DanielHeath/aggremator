package main

/*
TODO:
  * PA -> title has escaped entities
	* write 'go generate' based curl-to-sample thingy instead of manually updating
	* better way to test single item parsing than re-running everything
	* I'd prefer to extract the feeds to a config file so this could be re-usable
	* amazingsuperpowers had a youtube link in their feed, came thru as a busted thingy.
*/

import (
	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds"
	"github.com/danielheath/aggremator/feeds/alicegrove"
	"github.com/danielheath/aggremator/feeds/amazingsuperpowers"
	"github.com/danielheath/aggremator/feeds/dilbert"
	"github.com/danielheath/aggremator/feeds/orderofthestick"
	"github.com/danielheath/aggremator/feeds/pennyarcade"
	"github.com/danielheath/aggremator/feeds/questionablecontent"
	"github.com/danielheath/aggremator/feeds/smbc"
	"github.com/danielheath/aggremator/feeds/xkcd"
	"github.com/danielheath/aggremator/maildir"
	"github.com/danielheath/aggremator/pastentries"
	"github.com/go-gomail/gomail"
	"os/user"
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
	orderofthestick.Feed,
}
var currentFeeds = allFeeds // []feeds.Feed{
// 	orderofthestick.Feed,
// }

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

	for _, feed := range currentFeeds {
		var doc *rss.Feed
		if debug {
			doc, err = rss.Parse([]byte(feed.Sample()))
		} else {
			doc, err = rss.Fetch(feed.Url())
		}
		die(err)
		for _, item := range doc.Items {
			// TODO: One goroutine per feed
			// Have we seen this feed entry before?
			// TODO: pastEntries should be a smarter type, incorporating CleanId, not just a map
			if _, ok := pastEntries[maildir.CleanId(item.Link+item.ID)]; !ok {
				pastEntries[maildir.CleanId(item.Link+item.ID)] = true
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss@example.org")
				msg.SetHeader("To", "rss@example.org")
				err := feed.Serialize(*item, msg)
				die(err, item, "\n", item.Content)
				sender := maildir.Mailer(
					maildirPath +
						"/INBOX.Feeds." +
						feed.Category() +
						"/new/" +
						maildir.CleanId(item.ID),
				)

				m := gomail.NewMailer("localhost", "dummy", "dummy", 9002, gomail.SetSendMail(sender))
				die(m.Send(msg))

				if !debug {
					die(pastEntries.Write(pastEntriesFile)) // Update past entries after each message.
				}
			}
		}
	}
}

func die(err error, context ...interface{}) {
	if err != nil {
		fmt.Println(context)
		panic(err)
	}
}
