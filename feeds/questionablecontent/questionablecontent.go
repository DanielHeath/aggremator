package questionablecontent

// TODO: Comics often appear in the RSS feed some time before the file is available on the server.

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
