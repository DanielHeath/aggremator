package pennyarcade

// TODO: Start testing these properly using the sample.

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
	"regexp"
)

type Feed struct{}

func (f Feed) Url() string {
	return "https://penny-arcade.com/feed"
}
func (f Feed) Category() string {
	return "Comics.PennyArcade"
}
func (f Feed) Sample() string {
	return Sample
}

var (
	newsPostTitlePattern = regexp.MustCompile("News Post:")
	comicTitlePattern    = regexp.MustCompile("Comic:")
)

func (f Feed) Serialize(item rss.Item, msg *gomail.Message) error {
	msg.SetHeader("Subject", item.Title)

	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		return err
	}

	if newsPostTitlePattern.MatchString(item.Title) {
		post := doc.Find(".postBody .copy")
		return mail.AttachHtmlBody(msg, post, item.Link)
	} else if comicTitlePattern.MatchString(item.Title) {
		img := doc.Find("#comicFrame")
		return mail.AttachHtmlBody(msg, img, item.Link)
	} else {
		msg.AddAlternative("text/plain", item.Link+"\n\n"+item.Content)
		msg.SetBody("text/html", item.Content)
	}
	return nil
}
