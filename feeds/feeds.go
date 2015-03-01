package feeds

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
	"golang.org/x/net/html"
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

	fmt.Println("Attaching support the artist content && error message:")
	fmt.Println(doc.BeforeHtml(origError).AfterHtml(f.SupportTheArtist))

	selection := append(
		[]*html.Node{textToNode(origError)},
		doc.Nodes...,
	)
	// Can't see a tidier way to do this :/
	selection = append(selection, textToNode(f.SupportTheArtist))
	return mail.AttachHtmlBody(
		msg,
		*doc.Url,
		selection...,
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

	nodes := append(selection.Nodes, textToNode(f.SupportTheArtist))
	return mail.AttachHtmlBody(msg, *doc.Url, nodes...)
}

func textToNode(s string) *html.Node {
	b := bytes.NewBufferString("<div>" + s + "</div>")

	node, err := html.Parse(b)
	if err != nil {
		panic(fmt.Errorf(
			"Misconfiguration: Support URL did not parse (%s); html was %s",
			err.Error(),
			"<div>"+s+"</div>",
		))
	}
	return node
}
