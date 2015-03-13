package feeds

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
	"golang.org/x/net/html"
)

type Feed interface {
	Url() string
	Category() string
	Sample() string
	Serialize(rss.Item, *gomail.Message) error
}

type SelectorFunc func(*goquery.Document, rss.Item) ([]*html.Node, error)

func CssSelector(css string) SelectorFunc {
	return SelectorFunc(func(doc *goquery.Document, _ rss.Item) ([]*html.Node, error) {
		return doc.Find(css).Nodes, nil
	})
}

func MultiSelectorFunc(sf ...SelectorFunc) SelectorFunc {
	return SelectorFunc(func(doc *goquery.Document, item rss.Item) ([]*html.Node, error) {
		r := []*html.Node{}
		for _, s := range sf {
			nodes, err := s(doc, item)
			if err != nil {
				return r, err
			}
			r = append(r, nodes...)
		}
		return r, nil
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
	// Very rarely, a feed contains a bad link on purpose.
	// it's more useful to inline the feed content than fail in that case.
	// TODO: Only inline feed content when the <link> is an invalid URL.

	// Failed fetching the linked content; inline the feed content instead.
	origError := err.Error()

	// We're showing inline content from the feed; use the FeedUrl as the base url.
	u, err := url.Parse(f.FeedUrl)
	if err != nil {
		return fmt.Errorf("Selectorfeed FeedUrl is not a valid URL: %s", err)
	}

	reader := bytes.NewBufferString(item.Content)
	doc, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		return fmt.Errorf(
			"Failed to fetch linked content (%s), and failed falling back to using inline content (%s)",
			origError,
			err.Error(),
		)
	}
	fmt.Println(origError)
	fmt.Println("Attaching support the artist content && error message:")
	fmt.Println(f.FeedUrl)
	fmt.Println(item.Link)
	fmt.Println(item.Title)
	fmt.Println(doc.BeforeHtml(origError).AfterHtml(f.SupportTheArtist))

	selection := append(
		[]*html.Node{textToNode(origError)},
		doc.Nodes...,
	)
	// Can't see a tidier way to do this :/
	selection = append(selection, textToNode(f.SupportTheArtist))

	return mail.AttachHtmlBody(
		msg,
		*u,
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

	if len(selection) == 0 {
		return fmt.Errorf("No content found by selector function.")
	}

	nodes := append(selection, textToNode(f.SupportTheArtist))
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
