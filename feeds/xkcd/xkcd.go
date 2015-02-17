package xkcd

import (
	"bytes"
	"encoding/xml"
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
)

const Url = "http://xkcd.com/rss.xml"
const Category = "Comics.Xkcd"

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
	// msg.AddAlternative("text/plain", item.Link)
	return nil
}
