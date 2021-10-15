package parsers

import (
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Americanas struct{}

func (a *Americanas) ParseName(doc *html.Node) (string, error) {
	element, err := htmlquery.Query(doc, "//h1[1]")

	if err != nil {
		return "", err
	}

	name := htmlquery.InnerText(element)

	return name, nil
}

func (a *Americanas) ParsePrice(doc *html.Node) (float64, error) {
	element, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/main/div[2]/div[2]/div[1]/div[2]/div")

	if err != nil {
		return 0.0, err
	}

	priceAsString := htmlquery.InnerText(element)
	priceAsString = strings.ReplaceAll(priceAsString, "R$ ", "")
	priceAsString = strings.ReplaceAll(priceAsString, "%", "")
	priceAsString = strings.ReplaceAll(priceAsString, ".", "")
	priceAsString = strings.ReplaceAll(priceAsString, ",", ".")
	price, err := strconv.ParseFloat(priceAsString, 64)

	if err != nil {
		return 0.0, err
	}

	return price, nil
}
