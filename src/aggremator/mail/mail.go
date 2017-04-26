package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	gomail "gopkg.in/gomail.v1"
)

func GetBase64(url string) (string, error) {
	time.Sleep(time.Second) // Super hacky way to avoid slamming the server
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Error %s (%d) fetching %s", resp.Status, resp.StatusCode, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]

	return fmt.Sprintf(
		"data:%s;base64,%s",
		contentType,
		base64.StdEncoding.EncodeToString(body),
	), nil
}

func GetAttachment(url string, name string) (*gomail.File, error) {
	time.Sleep(time.Second) // Super hacky way to avoid slamming the server
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Error %s (%d) fetching %s", resp.Status, resp.StatusCode, url)
	}

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

// func InlineHtmlBody(msg *gomail.Message, baseUrl url.URL, nodes ...*html.Node) error {
// 	if len(nodes) <= 0 {
// 		return fmt.Errorf("AttachHtmlBody called with no body")
// 	}

// 	hasSrcAttr := cascadia.MustCompile("[src]")
// 	hasInlineStyle := cascadia.MustCompile("[style]")
// 	for _, topLevelNodes := range nodes {
// 		// Strip inline styles
// 		for _, styledNode := range hasInlineStyle.MatchAll(topLevelNodes) {
// 			setAttr(styledNode, "style", "")
// 		}

// 		for _, srcNode := range hasSrcAttr.MatchAll(topLevelNodes) {
// 			if src := getAttr(srcNode, "src"); src != "" {
// 				if srcset := getAttr(srcNode, "srcset"); srcset != "" {
// 					src = strings.Split(srcset, " ")[0]
// 				}
// 				src = rewriteSrc(src, baseUrl)
// 				// Add the <img> title attribute as literal text
// 				srcNode.Parent.InsertBefore(
// 					&html.Node{
// 						Type: html.TextNode,
// 						Data: getAttr(srcNode, "title"),
// 					},
// 					srcNode,
// 				)
// 				// Get the image
// 				img, err := GetBase64(src)
// 				if err != nil {
// 					return err
// 				}

// 				setAttr(srcNode, "orig-src", src)
// 				setAttr(srcNode, "src", img)
// 				setAttr(srcNode, "apple-inline", "yes")
// 			}
// 		}
// 	}

// 	richtext := msg.GetBodyWriter("text/html")
// 	richdoc := &html.Node{}
// 	richdoc.Type = html.DocumentNode
// 	for _, n := range nodes {
// 		// Get it out of its old document
// 		if n.Parent != nil {
// 			n.Parent.RemoveChild(n)
// 		}
// 		// Put it into a new one
// 		richdoc.AppendChild(n)
// 	}
// 	err := html.Render(richtext, richdoc)
// 	if err != nil {
// 		return err
// 	}

// 	plaintext := []string{"From " + baseUrl.String() + "\n"}
// 	for _, n := range nodes {
// 		plaintext = append(plaintext, renderPlainText(n))
// 	}

// 	msg.AddAlternative(
// 		"text/plain",
// 		strings.Join(plaintext, ""),
// 	)
// 	return nil
// }

// Woo, tonnes of duplicate code here.
func AttachHtmlBody(msg *gomail.Message, baseUrl url.URL, nodes ...*html.Node) error {
	if len(nodes) <= 0 {
		return fmt.Errorf("AttachHtmlBody called with no body")
	}

	attachmentCount := 0
	hasBody := cascadia.MustCompile("body")
	hasSrcAttr := cascadia.MustCompile("[src]")
	hasInlineStyle := cascadia.MustCompile("[style]")
	for _, topLevelNodes := range nodes {
		// Strip inline styles
		for _, styledNode := range hasInlineStyle.MatchAll(topLevelNodes) {
			setAttr(styledNode, "style", "")
		}

		for _, srcNode := range hasSrcAttr.MatchAll(topLevelNodes) {
			if src := getAttr(srcNode, "src"); src != "" {
				if srcset := getAttr(srcNode, "srcset"); srcset != "" {
					src = strings.Split(srcset, " ")[0]
				}
				src = rewriteSrc(src, baseUrl)
				// Add the <img> title attribute as literal text
				srcNode.Parent.InsertBefore(
					&html.Node{
						Type: html.TextNode,
						Data: getAttr(srcNode, "title"),
					},
					srcNode,
				)
				// Get the image
				img, err := GetAttachment(src, fmt.Sprintf("external_%d", attachmentCount))
				attachmentCount += 1
				if err != nil {
					return err
				}

				setAttr(srcNode, "orig-src", src)
				setAttr(srcNode, "src", "cid:"+img.Name)
				setAttr(srcNode, "apple-inline", "yes")
				msg.Attach(img)
			}
		}
	}

	richtext := msg.GetBodyWriter("text/html")

	richdoc, err := html.Parse(bytes.NewBufferString(`<html><head></head><body></body></html>`))
	for _, n := range nodes {
		if body := hasBody.MatchFirst(n); body != nil {
			n = body
			// Convert body to div
			n.Data = "div"
		}
		// Get it out of its old document
		if n.Parent != nil {
			n.Parent.RemoveChild(n)
		}
		// Put it into the new one
		richdoc.LastChild.LastChild.AppendChild(n)
	}

	err = html.Render(richtext, richdoc)
	if err != nil {
		return err
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
