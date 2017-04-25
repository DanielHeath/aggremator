package amazingsuperpowers

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://feeds.feedburner.com/amazingsuperpowers?format=xml",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".comicpane"),
		SupportTheArtist: "http://www.topatoco.com/merchant.mvc?Screen=CTGY&Store_Code=TO&Category_Code=ASP",
	}
)
