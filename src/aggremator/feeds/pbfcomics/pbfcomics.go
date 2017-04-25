package pbfcomics

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          `http://www.pbfcomics.com/feed/feed.xml`,
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector("#toptd"),
		SupportTheArtist: "http://www.pbfcomics.com/things/prints/",
	}
)
