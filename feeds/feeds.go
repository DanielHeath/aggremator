package feeds

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
	"strings"
)

type Feed interface {
	Url() string
	Category() string
	Sample() string
	Serialize(rss.Item, *gomail.Message) error
}

type SelectorFunc func(*goquery.Document, rss.Item) (*goquery.Selection, error)

func CssSelector(css string) SelectorFunc {
	return SelectorFunc(func(doc *goquery.Document, _ rss.Item) (*goquery.Selection, error) {
		return doc.Find(css), nil
	})
}

func ParentSelector(f SelectorFunc) SelectorFunc {
	return SelectorFunc(func(doc *goquery.Document, item rss.Item) (*goquery.Selection, error) {
		sel, err := f(doc, item)
		// fmt.Println("ASDFG")
		// fmt.Println(sel.Length())
		// fmt.Println(sel.Parent().Html())
		// fmt.Println("ASDFG")
		return sel.Parent(), err
	})
}

type SelectorFeed struct {
	FeedUrl          string
	FeedSample       string
	MailCategory     string
	SupportTheArtist string
	Selector         SelectorFunc
}

func (f SelectorFeed) Url() string {
	return f.FeedUrl
}
func (f SelectorFeed) Category() string {
	return f.MailCategory
}
func (f SelectorFeed) Sample() string {
	return f.FeedSample
}

func (f SelectorFeed) Serialize(item rss.Item, msg *gomail.Message) error {
	msg.SetHeader("Subject", item.Title)
	var doc *goquery.Document
	var err error
	err = f.serializeLink(item, msg)
	if err == nil { // All went well
		return nil
	}
	// Failed fetching the linked content; inline the feed content instead.
	origError := err.Error()
	reader := bytes.NewBufferString(item.Content)
	doc, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		return fmt.Errorf(
			"Failed to fetch linked content (%s), and failed falling back to using inline content (%s)",
			origError,
			err.Error(),
		)
	}
	return mail.AttachHtmlBody(
		msg,
		doc.BeforeHtml(origError).AfterHtml(f.SupportTheArtist),
		doc,
	)
}

func (f SelectorFeed) serializeLink(item rss.Item, msg *gomail.Message) error {
	var doc *goquery.Document
	var err error

	link := strings.Trim(item.Link, " \n")

	doc, err = goquery.NewDocument(link)
	if err != nil {
		return err
	}
	selection, err := f.Selector(doc, item)
	if err != nil {
		return err
	}

	if selection.Length() == 0 {
		return fmt.Errorf("No content found by selector function.")
	}
	selection = selection.AfterHtml(f.SupportTheArtist)
	return mail.AttachHtmlBody(msg, selection, doc)
}
