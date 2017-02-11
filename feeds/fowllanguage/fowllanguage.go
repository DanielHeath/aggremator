package fowllanguage

import (
	"github.com/danielheath/aggremator/feeds"
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
