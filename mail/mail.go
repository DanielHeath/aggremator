package mail

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.google.com/p/cascadia"
	"github.com/go-gomail/gomail"
	"golang.org/x/net/html"
)

func GetImg(url string) (string, error) {
	time.Sleep(time.Second) // Super hacky way to avoid slamming the server
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Error %s (%d) fetching %s", resp.Status, resp.StatusCode, url)
	}
	// setAttr(srcNode, "src", "data:"+contentType+";base64,"+)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	contentType := strings.Split(resp.Headers.Get("Content-Type"), ";")[0]

	return fmt.Sprintf(
		"data:%s;base64,%s",
		contentType,
		base64.StdEncoding.EncodeToString(body),
	), nil
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
				srcNode.Parent.InsertBefore(
					&html.Node{
						Type: html.TextNode,
						Data: getAttr(srcNode, "title"),
					},
					srcNode,
				)

				img, err := GetImg(src)
				if err != nil {
					return err
				}

				setAttr(srcNode, "orig-src", src)
				setAttr(srcNode, "src", img)
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
