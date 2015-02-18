package pennyarcade

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds/comic"
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
		html, err := post.Html()
		if err != nil {
			return err
		}
		text := post.Text()
		msg.SetBody("text/html", html)
		msg.AddAlternative("text/plain", item.Link+"\n\n"+text)
	} else if comicTitlePattern.MatchString(item.Title) {
		msg.AddAlternative("text/plain", item.Link)
		img := doc.Find("#comicFrame img")

		alt, _ := img.Attr("alt") // Not every img has an alt
		href, _ := img.Attr("src")

		if href == "" {
			return errors.New("No href found for " + item.Link) // every img has a src
		}

		return comic.Image{
			Title:    alt,
			Url:      href,
			Filename: "comic.jpg",
		}.AttachInline(msg)

	} else {
		msg.AddAlternative("text/plain", item.Content)
		msg.SetBody("text/html", item.Content)
	}
	return nil
}
