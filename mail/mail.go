package mail

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"net/http"
)

func GetImg(url string, name string) (*gomail.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &gomail.File{
		MimeType: resp.Header.Get("Content-Type"),
		Name:     name,
		Content:  body,
	}, nil
}

func AttachHtmlBody(msg *gomail.Message, s *goquery.Selection, link string) error {
	var err error
	if s.Length() == 0 {
		return fmt.Errorf("AttachHtmlBody called with no body")
	}
	externalResources := s.Find("[src]")
	resourceUrls := make([]*goquery.Selection, externalResources.Length())
	externalResources.Each(func(idx int, s *goquery.Selection) {
		resourceUrls[idx] = s
	})
	for idx, s := range resourceUrls {
		src, _ := s.Attr("src")
		if src != "" {
			img, err := GetImg(src, fmt.Sprintf("external_%d", idx))
			if err != nil {
				return err
			}
			msg.Embed(img)
			s.SetAttr("src", "cid:"+img.Name)
		}
	}
	html, err := s.Html()
	if err != nil {
		return err
	}

	msg.SetBody("text/html", html)

	msg.AddAlternative("text/plain", link+"\n\n"+s.Text())
	return nil
}
