package repositories

import (
	"encoding/json"
	"io/ioutil"
)

type Url struct {
	Url, Parser string
}

func (url Url) All() ([]Url, error) {
	var urls []Url

	urlsFile, err := ioutil.ReadFile("urls.json")
	urlsJson := (string(urlsFile))

	if err != nil {
		return urls, err
	}

	err = json.Unmarshal([]byte(urlsJson), &urls)

	if err != nil {
		return urls, err
	}

	return urls, nil
}
