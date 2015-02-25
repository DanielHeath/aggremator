package mail

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

// TODO: This only attaches the first selection item.
func AttachHtmlBody(msg *gomail.Message, s *goquery.Selection, doc *goquery.Document) error {
	var err error
	if s.Length() == 0 {
		return fmt.Errorf("AttachHtmlBody called with no body")
	}

	link := ""
	if doc.Url != nil {
		link = doc.Url.String()
	}

	externalResources := s.Find("[src]")
	resourceUrls := make([]*goquery.Selection, externalResources.Length())
	externalResources.Each(func(idx int, s *goquery.Selection) {
		resourceUrls[idx] = s
	})
	for idx, s := range resourceUrls {
		src, _ := s.Attr("src")

		if src != "" {
			if src[0] == '/' { // relative to root
				u, err := url.Parse(link)
				if err != nil {
					return fmt.Errorf("Looking for a domain for relative url '%s'; failed to parse document url (%s). TODO: Fallback to feed host?", link, err)
				}
				u.Path = src
				src = u.String()

			} else if src[:4] != "http" { // relative to page
				src = link[:strings.LastIndex(link, "/")] + "/" + src
			}
			img, err := GetImg(src, fmt.Sprintf("external_%d", idx))
			if err != nil {
				return err
			}
			msg.Embed(img)
			s.SetAttr("orig-src", src)
			s.SetAttr("src", "cid:"+img.Name)
		}
	}

	html := ""
	s.EachWithBreak(func(_ int, i *goquery.Selection) bool {
		var h string
		h, err = i.Html()
		if err != nil {
			return false
		}
		html += h
		return true
	})
	if err != nil {
		return err
	}

	msg.SetBody("text/html", html)

	msg.AddAlternative("text/plain", link+"\n\n"+s.Text())
	return nil
}
