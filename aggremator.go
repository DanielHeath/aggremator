package main

/*
TODO:
 * More than just XKCD :)
 * what belongs where?
   * feed-item -> mail serialization (with the feed)
*/

import (
	"bytes"
	"encoding/xml"
	"flag"
	"github.com/SlyMarbo/rss"
	// "github.com/danielheath/aggremator/feeds"
	"github.com/danielheath/aggremator/feeds/xkcd"
	"github.com/danielheath/aggremator/maildir"
	"github.com/danielheath/aggremator/pastentries"
	"github.com/go-gomail/gomail"
	"os/user"
)

var pastEntriesPath string
var xkcdFolder string

func init() {
	usr, err := user.Current()
	die(err)
	xkcdFolder = usr.HomeDir + "/.mail/fastmail/INBOX.Feeds.Comics.Xkcd/new/"
	flag.StringVar(
		&pastEntriesPath,
		"pastEntriesFile",
		usr.HomeDir+"/.aggremator/pastentries",
		"File to store which feed items have already been sync-ed",
	)
}

func main() {
	flag.Parse()

	pastEntriesFile := pastentries.File(pastEntriesPath)
	pastEntries, err := pastEntriesFile.Read()
	die(err)
	feed, err := rss.Fetch(xkcd.Url)
	die(err)
	for _, item := range feed.Items {
		if _, ok := pastEntries[maildir.CleanId(item.ID)]; !ok {
			pastEntries[maildir.CleanId(item.ID)] = true
			v := xkcd.XkcdImg{}
			err := xml.NewDecoder(bytes.NewBufferString(item.Content)).Decode(&v)
			die(err)
			msg := gomail.NewMessage()
			msg.SetHeader("From", "rss@example.org")
			msg.SetHeader("To", "rss@example.org")

			sender := maildir.Mailer(xkcdFolder + maildir.CleanId(item.ID))
			m := gomail.NewMailer("localhost", "dummy", "dummy", 9002, gomail.SetSendMail(sender))
			m.Send(msg)
		}
	}
	pastEntries.Write(pastEntriesFile)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
