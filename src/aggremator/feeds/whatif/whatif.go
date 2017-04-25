package whatif

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://what-if.xkcd.com/feed.atom",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector("article.entry"),
		SupportTheArtist: "<a href='http://store.xkcd.com/'>Support the artist!</a>",
	}
)
