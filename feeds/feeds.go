package feeds

import (
	"github.com/SlyMarbo/rss"
	"github.com/go-gomail/gomail"
)

type Feed interface {
	Url() string
	Category() string
	Sample() string
	Serialize(rss.Item, *gomail.Message) error
}
