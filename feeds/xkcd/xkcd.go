package xkcd

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "https://xkcd.com/rss.xml",
		FeedSample:       Sample,
		MailCategory:     "Comics.Xkcd",
		Selector:         feeds.CssSelector("#comic"),
		SupportTheArtist: "<a href='http://store.xkcd.com/'>Support the artist!</a>",
	}
)
