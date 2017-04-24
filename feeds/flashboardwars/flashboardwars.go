package flashboardwars

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl: "http://theagevsheraldsun.tumblr.com/rss",
		// FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector("#stat-articles > article"),
		SupportTheArtist: "",
	}
)
