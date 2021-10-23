package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/pedropazello/price-time-machine/src/models"
	"github.com/pedropazello/price-time-machine/src/observers"
	"github.com/pedropazello/price-time-machine/src/parsers"
	"golang.org/x/net/html"
)

func GetOffer(url string, parser parsers.Parser, observers []observers.Observer) error {
	urlError := errors.New("URL failed: " + url).Error()

	doc, err := htmlDocByUrl(url)

	if err != nil {
		return errors.New(urlError + " - " + err.Error())
	}

	offer := models.Offer{}
	offer.ProductName, err = parser.ParseName(doc)

	if err != nil {
		return errors.New(urlError + " - " + err.Error())
	}

	offer.Price, err = parser.ParsePrice(doc)

	if err != nil {
		return errors.New(urlError + " - " + err.Error())
	}

	for i := range observers {
		observers[i].Execute(offer)
	}

	return nil
}

func htmlDocByUrl(url string) (*html.Node, error) {
	body, err := getUrlResponseBody(url)

	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getUrlResponseBody(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0"},
		"Accept-Language": []string{"pt-BR,pt;q=0.8,en-US;q=0.5,en;q=0.3"},
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
