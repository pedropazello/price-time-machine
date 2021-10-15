package main

import (
	"github.com/pedropazello/price-time-machine/src/observers"
	"github.com/pedropazello/price-time-machine/src/parsers"
	"github.com/pedropazello/price-time-machine/src/services"
)

func main() {
	americanasUrl := "https://www.americanas.com.br/produto/2779138678?pfm_carac=bicicleta-caloi-vulcan&pfm_page=search&pfm_pos=grid&pfm_type=search_page&offerId=607862e60c070442666ca14f&buyboxToken=smartbuybox-acom-v2-addaba6c-8376-4596-8176-344fc36d2414-2021-10-12%2011%3A01%3A00-0300&cor=BRANCO&tamanho=17"
	crawler := parsers.Americanas{}
	err := services.GetOffer(americanasUrl, &crawler, []observers.Observer{&observers.OfferPrinter{}})

	if err != nil {
		panic(err)
	}

	mercadoLivreUrl := "https://produto.mercadolivre.com.br/MLB-1095601855-bicicleta-viking-tuff-21-velocidadesfreio-a-disco-_JM#reco_item_pos=2&reco_backend=machinalis-homes-pdp-boos&reco_backend_type=function&reco_client=home_navigation-trend-recommendations&reco_id=c3651eb8-1f61-4a90-b419-fe481231f656&c_id=/home/navigation-trends-recommendations/element&c_element_order=3&c_uid=2e7f9525-43a4-466b-88fa-50b4d4bbf2b1"
	crawlerMercadoLivre := parsers.MercadoLivre{}
	err = services.GetOffer(mercadoLivreUrl, &crawlerMercadoLivre, []observers.Observer{&observers.OfferPrinter{}})

	if err != nil {
		panic(err)
	}
}
