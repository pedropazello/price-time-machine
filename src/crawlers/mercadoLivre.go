package crawlers

import "github.com/pedropazello/price-time-machine/src/models"

type MercadoLivre struct{}

func (m *MercadoLivre) GetOfferFor(url string) models.Offer {
	offer := models.Offer{}
	return offer
}
