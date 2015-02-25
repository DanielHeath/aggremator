package comic

// TODO: Merge into 'feeds'
import (
	"github.com/danielheath/aggremator/mail"
	"github.com/go-gomail/gomail"
)

type Image struct {
	Title    string
	Url      string
	Filename string
}

func (i Image) AttachInline(msg *gomail.Message) error {
	msg.SetHeader("Subject", i.Title)

	img, err := mail.GetImg(i.Url, i.Filename)
	if err != nil {
		return err
	}

	msg.Embed(img)

	msg.SetBody(
		"text/html",
		`<img
      src="cid:`+img.Name+`"
      alt="`+i.Title+`"
      title="`+i.Title+`"
     ></img>
    `,
	)
	return nil
}
