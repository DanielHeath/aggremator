package codelesscode

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://thecodelesscode.com/rss",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".titles, .contenttext"),
		SupportTheArtist: "http://thecodelesscode.com",
	}
)
