package maildir

import (
	"io/ioutil"
	"net/smtp"
	"os"
	"regexp"
)

var cleanId = regexp.MustCompile("[^\\w]")

func CleanId(id string) string {
	return string(cleanId.ReplaceAll([]byte(id), []byte{}))
}

func Mailer(path string) func(string, smtp.Auth, string, []string, []byte) error {
	return func(_ string, _ smtp.Auth, from string, to []string, msg []byte) error {
		// todo: ensure directory exists
		return ioutil.WriteFile(path, msg, os.ModePerm)
	}
}
