package dilbert

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds/comic"
	"github.com/go-gomail/gomail"
)

type Feed struct{}

func (f Feed) Url() string {
	return "http://feed.dilbert.com/dilbert/daily_strip"
}
func (f Feed) Category() string {
	return "Comics.Dilbert"
}
func (f Feed) Sample() string {
	return Sample
}

func (f Feed) Serialize(item rss.Item, msg *gomail.Message) error {
	msg.SetHeader("Subject", item.Title)
	msg.AddAlternative("text/plain", item.Link)

	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		return err
	}

	c := doc.Find(".img-comic-container [alt]")
	href, _ := c.Attr("src")
	alt, _ := c.Attr("alt")

	// href, _ := doc.Find(`meta[property="og:image"]`).Attr("content")

	if href == "" {
		return errors.New("No href found for " + item.Link) // every img has a src
	}

	return comic.Image{
		Title:    alt,
		Url:      href,
		Filename: "comic.gif",
	}.AttachInline(msg)

	return nil
}
