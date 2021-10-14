package observers

import (
	"fmt"

	"github.com/pedropazello/price-time-machine/src/models"
)

type OfferPrinter struct{}

func (o *OfferPrinter) Execute(offer models.Offer) {
	fmt.Println(offer)
}
