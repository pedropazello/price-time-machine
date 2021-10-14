package services

import (
	"github.com/pedropazello/price-time-machine/src/crawlers"
	"github.com/pedropazello/price-time-machine/src/observers"
)

func GetOffer(url string, crawler crawlers.Crawler, observers []observers.Observer) {
	offer := crawler.GetOfferFor(url)

	for i := range observers {
		observers[i].Execute(offer)
	}

}
