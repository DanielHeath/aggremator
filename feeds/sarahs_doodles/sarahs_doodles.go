package sarahs_doodles

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "https://tapastic.com/rss/series/2007",
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".ep-contents"),
		SupportTheArtist: "",
	}
)
