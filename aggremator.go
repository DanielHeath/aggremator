package main

/*
TODO:
  * Save *all* inline images from html pages
  * PA -> title has escaped entities
	* Start testing some things
	* goquery for every link attr?
*/

import (
	"flag"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds"
	"github.com/danielheath/aggremator/feeds/amazingsuperpowers"
	"github.com/danielheath/aggremator/feeds/dilbert"
	"github.com/danielheath/aggremator/feeds/pennyarcade"
	"github.com/danielheath/aggremator/feeds/xkcd"
	"github.com/danielheath/aggremator/maildir"
	"github.com/danielheath/aggremator/pastentries"
	"github.com/go-gomail/gomail"
	"os/user"
)

var pastEntriesPath string
var homedir string
var debug bool

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
		true,
		"debug mode",
	)
}

var allFeeds = []feeds.Feed{
	xkcd.Feed{},
	pennyarcade.Feed{},
	dilbert.Feed{},
	amazingsuperpowers.Feed{},
}

func main() {
	flag.Parse()

	pastEntriesFile := pastentries.File(pastEntriesPath)
	pastEntries, err := pastEntriesFile.Read()
	if debug {
		pastEntries = pastentries.PastEntries{}
	}
	die(err)

	for _, feed := range allFeeds {
		var doc *rss.Feed
		if debug {
			doc, err = rss.Parse([]byte(feed.Sample()))
		} else {
			doc, err = rss.Fetch(feed.Url())
		}
		die(err)
		for _, item := range doc.Items {
			// Have we seen this feed entry before?
			// TODO: pastEntries should be a smarter type, incorporating CleanId, not just a map
			if _, ok := pastEntries[maildir.CleanId(item.Link+item.ID)]; !ok {
				pastEntries[maildir.CleanId(item.Link+item.ID)] = true
				msg := gomail.NewMessage()
				msg.SetHeader("From", "rss@example.org")
				msg.SetHeader("To", "rss@example.org")
				err := feed.Serialize(*item, msg)
				die(err)
				sender := maildir.Mailer(
					homedir +
						"/.mail/fastmail/INBOX.Feeds." +
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

func die(err error) {
	if err != nil {
		panic(err)
	}
}
