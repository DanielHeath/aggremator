package amazingsuperpowers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
)

type Feed struct{}

func (f Feed) Url() string {
	return "http://feeds.feedburner.com/amazingsuperpowers?format=xml"
}
func (f Feed) Category() string {
	return "Comics.AmazingSuperPowers"
}
func (f Feed) Sample() string {
	return Sample
}

func (f Feed) Serialize(item rss.Item, msg *gomail.Message) error {
	msg.SetHeader("Subject", item.Title)
	msg.AddAlternative("text/plain", item.Link)

	c, err := comic(item)
	if err != nil {
		return err
	}
	return mail.AttachHtmlBody(msg, c, item.Link)
}

func comic(item rss.Item) (*goquery.Selection, error) {
	doc, err := goquery.NewDocument(item.Link)

	if err != nil {
		return nil, err
	}
	return sliceComic(doc)
}

func sliceComic(f *goquery.Document) (*goquery.Selection, error) {
	return f.Find(".comicpane"), nil
}
