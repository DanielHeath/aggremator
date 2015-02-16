package main

/*
TODO:
 * More than just XKCD :)
*/

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"os/user"
	"regexp"
)

const xkcd = "http://xkcd.com/rss.xml"

var cleanId = regexp.MustCompile("[^\\w]")
var pastEntriesFile string
var xkcdFolder string

func init() {
	usr, err := user.Current()
	die(err)
	xkcdFolder = usr.HomeDir + "/.mail/fastmail/INBOX.Feeds.Comics.Xkcd/new/"
	flag.StringVar(
		&pastEntriesFile,
		"pastEntriesFile",
		usr.HomeDir+"/.aggremator/pastentries",
		"File to store which feed items have already been sync-ed",
	)
}

type XkcdImg struct {
	Src   string `xml:"src,attr"`
	Title string `xml:"title,attr"`
	Alt   string `xml:"alt,attr"`
}

func (xi XkcdImg) ToXml() string {
	b := bytes.Buffer{}
	die(xml.NewEncoder(&b).Encode(xi))
	return b.String()
}

func getImg(url string) gomail.File {
	resp, err := http.Get(url)
	die(err)
	body, err := ioutil.ReadAll(resp.Body)
	die(err)
	defer resp.Body.Close()
	return gomail.File{
		MimeType: resp.Header.Get("Content-Type"),
		Name:     "image.png",
		Content:  body,
	}
}

func WriteToMailDir(path string) gomail.SendMailFunc {
	return func(_ string, _ smtp.Auth, from string, to []string, msg []byte) error {
		return ioutil.WriteFile(path, msg, os.ModePerm)
	}
}
func CleanId(id string) string {
	return string(cleanId.ReplaceAll([]byte(id), []byte{}))
}

func readPastEntries(path string) map[string]bool {
	pastEntries := make(map[string]bool)

	file, err := os.Open(path)
	die(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pastEntries[scanner.Text()] = true
	}
	die(scanner.Err())
	return pastEntries
}

func writePastEntries(path string, entries map[string]bool) {
	data := bytes.Buffer{}
	for k, _ := range entries {
		fmt.Fprintln(&data, k)
	}
	die(ioutil.WriteFile(path, data.Bytes(), os.ModePerm))
}

func main() {
	flag.Parse()
	pastEntries := readPastEntries(pastEntriesFile)

	feed, err := rss.Fetch(xkcd)
	die(err)
	for _, item := range feed.Items {
		if _, ok := pastEntries[CleanId(item.ID)]; !ok {
			pastEntries[CleanId(item.ID)] = true
			v := XkcdImg{}
			err := xml.NewDecoder(bytes.NewBufferString(item.Content)).Decode(&v)
			die(err)
			msg := gomail.NewMessage()
			msg.SetHeader("From", "rss@example.org")
			msg.SetHeader("To", "rss@example.org")
			msg.SetHeader("Subject", v.Title)

			img := getImg(v.Src)
			msg.Embed(&img)
			v.Src = "cid:" + img.Name
			msg.SetBody("text/html", v.ToXml())
			msg.AddAlternative("text/plain", item.Link)
			sender := WriteToMailDir(xkcdFolder + CleanId(item.ID))
			m := gomail.NewMailer("localhost", "dummy", "dummy", 9002, gomail.SetSendMail(sender))
			m.Send(msg)
		}
	}
	writePastEntries(pastEntriesFile, pastEntries)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
