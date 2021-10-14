package crawlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/pedropazello/price-time-machine/src/models"
	"golang.org/x/net/html"
)

type Americanas struct {
}

func (a *Americanas) GetOfferFor(url string) models.Offer {
	doc := htmlDocByUrl(url)

	offer := models.Offer{}
	offer.ProductName = parseName(doc)
	offer.Url = url
	offer.Price = parsePrice(doc)
	return offer
}

func htmlDocByUrl(url string) *html.Node {
	body := getUrlResponseBody(url)

	doc, err := htmlquery.Parse(strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	return doc
}

func getUrlResponseBody(url string) string {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Host":            []string{"www.americanas.com.br"},
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0"},
		"Accept-Language": []string{"pt-BR,pt;q=0.8,en-US;q=0.5,en;q=0.3"},
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func parseName(doc *html.Node) string {
	element, err := htmlquery.Query(doc, "//h1[1]")

	if err != nil {
		panic(err)
	}

	name := htmlquery.InnerText(element)

	return name
}

func parsePrice(doc *html.Node) float64 {
	element, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/main/div[2]/div[2]/div[1]/div[2]/div")

	if err != nil {
		panic(err)
	}

	priceAsString := htmlquery.InnerText(element)
	priceAsString = strings.ReplaceAll(priceAsString, "R$ ", "")
	priceAsString = strings.ReplaceAll(priceAsString, "%", "")
	priceAsString = strings.ReplaceAll(priceAsString, ".", "")
	priceAsString = strings.ReplaceAll(priceAsString, ",", ".")
	price, err := strconv.ParseFloat(priceAsString, 64)

	if err != nil {
		panic(err)
	}

	return price
}
