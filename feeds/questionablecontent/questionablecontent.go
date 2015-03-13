package questionablecontent

import (
	"github.com/danielheath/aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl:          "http://www.questionablecontent.net/QCRSS.xml",
		FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector("#comic img"),
		SupportTheArtist: "http://www.topatoco.com/merchant.mvc?Screen=CTGY&Store_Code=TO&Category_Code=QC",
	}
)
