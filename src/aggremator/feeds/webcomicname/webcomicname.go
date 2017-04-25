package webcomicname

import (
	"aggremator/feeds"
)

var (
	Feed = feeds.SelectorFeed{
		FeedUrl: "http://webcomicname.com/rss",
		// FeedSample:       Sample,
		MailCategory:     "Comics",
		Selector:         feeds.CssSelector("#posts .post-content"),
		SupportTheArtist: "",
	}
)
