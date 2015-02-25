package dilbert

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://feed.dilbert.com/dilbert/daily_strip",
		FeedSample:       Sample,
		MailCategory:     "Comics.Dilbert",
		Selector:         feeds.CssSelector(".img-comic-container"),
		SupportTheArtist: "http://thedilbertstore.com/",
	}
)
