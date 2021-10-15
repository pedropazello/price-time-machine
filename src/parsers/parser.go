package parsers

import "golang.org/x/net/html"

type Parser interface {
	ParseName(doc *html.Node) (string, error)
	ParsePrice(doc *html.Node) (float64, error)
}
