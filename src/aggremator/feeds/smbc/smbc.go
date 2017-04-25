package smbc

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:      "http://www.smbc-comics.com/rss.php",
		FeedSample:   Sample,
		MailCategory: "Comics",
		Selector: feeds.MultiSelectorFunc(
			feeds.CssSelector("#cc-comic"),
			feeds.CssSelector("#aftercomic"),
		),
		SupportTheArtist: "http://smbc.myshopify.com and https://www.patreon.com/ZachWeinersmith",
	}
)
