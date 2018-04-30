package goodbear

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl: "https://goodbearcomics.com/feed/",
		// FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".article-content"),
		SupportTheArtist: "",
	}
)
