package pennyarcade

import (
	"fmt"
	"regexp"

	"aggremator/feeds"

	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"golang.org/x/net/html"
)

var contentSelector = feeds.CssSelector(".postBody .copy, #comicFrame")
var titleSelector = feeds.CssSelector("head title")

var (
	newsPostTitlePattern = regexp.MustCompile("News Post:")
	comicTitlePattern    = regexp.MustCompile("Comic:")

	Feed = feeds.SelectorFeed{
		FeedUrl:          "https://penny-arcade.com/feed",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		SupportTheArtist: "http://store.penny-arcade.com/",
		Selector: feeds.SelectorFunc(func(doc *goquery.Document, item rss.Item) ([]*html.Node, error) {
			titles, err := titleSelector(doc, item)
			if err != nil {
				return nil, err
			}

			for _, title := range titles {
				if title != nil {
					if title.Data == "Penny Arcade - 404" {
						return nil, fmt.Errorf("Content isn't up yet")
					}
				}
			}

			content, err := contentSelector(doc, item)
			if err != nil {
				return nil, err
			}
			if len(content) > 1 {
				return nil, fmt.Errorf("Newspost isn't up yet")
			}
			return content, nil
		}),
	}
)
