package xkcd

// tOdO: use 'comic' package to do images instead of XML bashing
import (
	"bytes"
	"encoding/xml"
	"github.com/SlyMarbo/rss"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
)

type Feed struct{}

func (f Feed) Url() string {
	return "https://xkcd.com/rss.xml"
}
func (f Feed) Category() string {
	return "Comics.Xkcd"
}
func (f Feed) Sample() string {
	return Sample
}

func (f Feed) Serialize(item rss.Item, msg *gomail.Message) error {
	v := XkcdImg{}
	err := xml.NewDecoder(bytes.NewBufferString(item.Content)).Decode(&v)
	if err != nil {
		return err
	}
	msg.AddAlternative("text/plain", item.Link)
	return v.Serialize(msg)
}

type XkcdImg struct {
	Src   string `xml:"src,attr"`
	Title string `xml:"title,attr"`
	Alt   string `xml:"alt,attr"`
}

func (xi XkcdImg) xml() (string, error) {
	b := bytes.Buffer{}
	err := xml.NewEncoder(&b).Encode(xi)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (xi XkcdImg) Serialize(msg *gomail.Message) error {
	msg.SetHeader("Subject", xi.Title)

	img, err := mail.GetImg(xi.Src, "xkcd.png")
	if err != nil {
		return err
	}
	msg.Embed(img)
	xi.Src = "cid:" + img.Name
	body, err := xi.xml()
	if err != nil {
		return err
	}
	msg.SetBody("text/html", body)
	return nil
}
