package maildir

import (
	"io/ioutil"
	"net/smtp"
	"os"
	"path"
	"regexp"
)

var cleanId = regexp.MustCompile("[^\\w]")

func CleanId(id string) string {
	return string(cleanId.ReplaceAll([]byte(id), []byte{}))
}

func Mailer(fpath string) func(string, smtp.Auth, string, []string, []byte) error {
	return func(_ string, _ smtp.Auth, from string, to []string, msg []byte) error {
		_ = os.MkdirAll(path.Dir(path.Dir(fpath))+"/tmp", os.ModePerm)
		_ = os.MkdirAll(path.Dir(path.Dir(fpath))+"/cur", os.ModePerm)
		_ = os.MkdirAll(path.Dir(path.Dir(fpath))+"/new", os.ModePerm)
		return ioutil.WriteFile(fpath, msg, os.ModePerm)
	}
}
