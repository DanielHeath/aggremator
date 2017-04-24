package questionablecontent

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds"
	"golang.org/x/net/html"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:      "http://www.questionablecontent.net/QCRSS.xml",
		FeedSample:   Sample,
		MailCategory: "Comics",
		Selector: feeds.SelectorFunc(func(doc *goquery.Document, item rss.Item) ([]*html.Node, error) {
			u, err := url.Parse(item.Link)
			if err != nil {
				return nil, err
			}
			if u.Host != "questionablecontent.net" {
				return feeds.RawItem(doc, item)
			}
			return feeds.CssSelector("img#strip")(doc, item)
		}),
		SupportTheArtist: "http://www.topatoco.com/merchant.mvc?Screen=CTGY&Store_Code=TO&Category_Code=QC",
	}
)
