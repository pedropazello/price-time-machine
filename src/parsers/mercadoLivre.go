package parsers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type MercadoLivre struct{}

func (m *MercadoLivre) ParseName(doc *html.Node) (string, error) {
	element, err := htmlquery.Query(doc, "/html/body/main/div/div[3]/div/div[1]/div/div[1]/div/div[1]/div/div[2]/h1")

	if element == nil {
		return "", errors.New("name not found")
	}

	if err != nil {
		return "", err
	}

	name := htmlquery.InnerText(element)

	return name, nil
}

func (m *MercadoLivre) ParsePrice(doc *html.Node) (float64, error) {
	element, err := htmlquery.Query(doc, "/html/body/main/div/div[3]/div/div[1]/div/div[1]/div/div[2]/div/div[1]/span/span[2]/span[2]")

	if element == nil {
		return 0.0, errors.New("price not found")
	}

	if err != nil {
		return 0.0, err
	}

	priceAsString := htmlquery.InnerText(element)
	priceAsString = strings.ReplaceAll(priceAsString, ".", "")
	priceAsString = strings.ReplaceAll(priceAsString, ",", ".")
	price, err := strconv.ParseFloat(priceAsString, 64)

	if err != nil {
		return 0.0, err
	}

	return price, nil
}
