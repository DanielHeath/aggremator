package smbc

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://feeds.feedburner.com/smbc-comics/PvLb",
		FeedSample:       Sample,
		MailCategory:     "Comics.SMBC",
		Selector:         feeds.CssSelector("#comicimage, #aftercomic"),
		SupportTheArtist: "http://smbc.myshopify.com and https://www.patreon.com/ZachWeinersmith",
	}
)
