package observers

import "github.com/pedropazello/price-time-machine/src/models"

type Observer interface {
	Execute(models.Offer) error
}
