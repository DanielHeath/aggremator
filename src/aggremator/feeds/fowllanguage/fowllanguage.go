package fowllanguage

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl: "http://www.fowllanguagecomics.com/feed/",
		// FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".main-img"),
		SupportTheArtist: "",
	}
)
