package dailywtf

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://syndication.thedailywtf.com/TheDailyWtf",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector(".article-body"),
		SupportTheArtist: "",
	}
)
