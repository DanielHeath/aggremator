package pennyarcade

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/feeds"
	"regexp"
)

var (
	newsPostTitlePattern = regexp.MustCompile("News Post:")
	comicTitlePattern    = regexp.MustCompile("Comic:")

	Feed = feeds.SelectorFeed{
		FeedUrl:          "https://penny-arcade.com/feed",
		FeedSample:       Sample,
		MailCategory:     "Comics.PennyArcade",
		SupportTheArtist: "http://store.penny-arcade.com/",
		Selector: feeds.SelectorFunc(func(doc *goquery.Document, item rss.Item) (*goquery.Selection, error) {
			if newsPostTitlePattern.MatchString(item.Title) {
				return feeds.CssSelector(".postBody .copy")(doc, item)
			} else if comicTitlePattern.MatchString(item.Title) {
				return feeds.CssSelector("#comicFrame")(doc, item)
			} else {
				return nil, errors.New("Feed item did not match a known format")
			}
		}),
	}
)
