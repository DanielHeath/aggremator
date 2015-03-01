package pennyarcade

import (
	"github.com/danielheath/aggremator/feeds"
	"regexp"
)

var (
	newsPostTitlePattern = regexp.MustCompile("News Post:")
	comicTitlePattern    = regexp.MustCompile("Comic:")

	Feed = feeds.SelectorFeed{
		FeedUrl:          "https://penny-arcade.com/feed",
		FeedSample:       Sample,
		MailCategory:     "Comics.PennyArcade",
		SupportTheArtist: "http://store.penny-arcade.com/",
		Selector:         feeds.CssSelector(".postBody .copy, #comicFrame"),
	}
)
