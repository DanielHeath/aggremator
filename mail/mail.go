package mail

import (
	"code.google.com/p/cascadia"
	"fmt"
	"github.com/go-gomail/gomail"
	"golang.org/x/net/html"
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

func replaceSrcWithAttachments(msg *gomail.Message, n *html.Node) {
	for _, a := range n.Attr {
		if a.Key == "src" {
			a.Val = ""
		}
	}
}

func getAttr(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

// Does not handle multiple copies of the same attr.
func setAttr(n *html.Node, attr, val string) {

	for i, a := range n.Attr {
		if a.Key == attr {
			n.Attr[i].Val = val
			return
		}
	}
	n.Attr = append(n.Attr, html.Attribute{Key: attr, Val: val})
	return
}

func rewriteSrc(src string, context url.URL) string {
	ctx, _ := context.Parse(src)
	return ctx.String()
}

func renderPlainText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	alt := getAttr(n, "alt")
	src := getAttr(n, "src")
	title := getAttr(n, "title")
	if title == alt {
		title = ""
	}
	if src != "" {
		src = "[" + src + "]"
	}

	results := []string{src}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		results = append(results, renderPlainText(child))
	}
	results = append(results, alt, title)
	return strings.Join(results, " ")
}

func AttachHtmlBody(msg *gomail.Message, baseUrl url.URL, nodes ...*html.Node) error {
	attachmentCount := 0
	if len(nodes) <= 0 {
		return fmt.Errorf("AttachHtmlBody called with no body")
	}

	hasSrcAttr := cascadia.MustCompile("[src]")
	hasInlineStyle := cascadia.MustCompile("[style]")
	for _, topLevelNodes := range nodes {
		// Strip inline styles
		for _, styledNode := range hasInlineStyle.MatchAll(topLevelNodes) {
			setAttr(styledNode, "style", "")
		}

		for _, srcNode := range hasSrcAttr.MatchAll(topLevelNodes) {
			if src := getAttr(srcNode, "src"); src != "" {
				src = rewriteSrc(src, baseUrl)
				img, err := GetImg(src, fmt.Sprintf("external_%d", attachmentCount))
				attachmentCount += 1
				if err != nil {
					return err
				}
				setAttr(srcNode, "orig-src", src)
				setAttr(srcNode, "src", "cid:"+img.Name)
				msg.Embed(img)
			}
		}
	}

	richtext := msg.GetBodyWriter("text/html")
	for _, n := range nodes {
		err := html.Render(richtext, n)
		if err != nil {
			return err
		}
	}

	plaintext := []string{"From " + baseUrl.String() + "\n"}
	for _, n := range nodes {
		plaintext = append(plaintext, renderPlainText(n))
	}

	msg.AddAlternative(
		"text/plain",
		strings.Join(plaintext, ""),
	)
	return nil
}
