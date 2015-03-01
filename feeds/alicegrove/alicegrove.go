package alicegrove

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://www.alicegrove.com/rss",
		FeedSample:       Sample,
		MailCategory:     "Comics.AliceGrove",
		Selector:         feeds.CssSelector(".photo-hires-item, .caption"),
		SupportTheArtist: "http://www.topatoco.com/merchant.mvc?Screen=CTGY&Store_Code=TO&Category_Code=QC",
	}
)
