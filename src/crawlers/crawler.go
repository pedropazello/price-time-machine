package crawlers

import "github.com/pedropazello/price-time-machine/src/models"

type Crawler interface {
	GetOfferFor(url string) models.Offer
}
