package smbc

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:      "http://feeds.feedburner.com/smbc-comics/PvLb",
		FeedSample:   Sample,
		MailCategory: "Comics",
		Selector: feeds.MultiSelectorFunc(
			feeds.CssSelector("#comicimage,#comicbody"),
			feeds.CssSelector("#aftercomic"),
		),
		SupportTheArtist: "http://smbc.myshopify.com and https://www.patreon.com/ZachWeinersmith",
	}
)
