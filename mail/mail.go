package mail

import (
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"net/http"
)

func GetImg(url string, name string) (*gomail.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &gomail.File{
		MimeType: resp.Header.Get("Content-Type"),
		Name:     name,
		Content:  body,
	}, nil
}
