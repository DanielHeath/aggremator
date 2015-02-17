package feeds

import (
	"github.com/go-gomail/gomail"
)

type Feed interface {
	Url() string
	Fetch() ([]FeedItem, error)
}

type FeedItem interface {
	Serialize(*gomail.Message) error
}
