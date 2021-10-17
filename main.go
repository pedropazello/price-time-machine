package main

import (
	"fmt"

	"github.com/pedropazello/price-time-machine/src/observers"
	"github.com/pedropazello/price-time-machine/src/parsers"
	"github.com/pedropazello/price-time-machine/src/repositories"
	"github.com/pedropazello/price-time-machine/src/services"
)

func main() {
	urls, err := repositories.Url{}.All()

	if err != nil {
		fmt.Println(err)
	}

	actions := []observers.Observer{&observers.OfferPrinter{}, &observers.OfferSQSSender{}}

	for _, site := range urls {
		var crawler parsers.Parser

		switch site.Parser {
		case "Americanas":
			crawler = &parsers.Americanas{}
		case "MercadoLivre":
			crawler = &parsers.MercadoLivre{}
		}

		err := services.GetOffer(site.Url, crawler, actions)

		if err != nil {
			fmt.Println(err)
		}
	}
}
