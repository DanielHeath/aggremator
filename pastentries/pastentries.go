package pastentries

// TODO: Extract interface, add postgres storage
import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type File string

type PastEntries map[string]bool

func (path File) Read() (PastEntries, error) {
	pastEntries := make(map[string]bool)
	file, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pastEntries[scanner.Text()] = true
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return pastEntries, nil
}

func (entries PastEntries) Write(path File) error {
	data := bytes.Buffer{}
	for k, _ := range entries {
		fmt.Fprintln(&data, k)
	}
	return ioutil.WriteFile(string(path), data.Bytes(), os.ModePerm)
}
