package feeds

import (
	"bytes"
	"fmt"
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
	return f.serializeLink(item, msg)
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
