package main

import (
	"fmt"

	"github.com/pedropazello/price-time-machine/src/crawlers"
)

func main() {
	url := "https://www.americanas.com.br/produto/2779138678?pfm_carac=bicicleta-caloi-vulcan&pfm_page=search&pfm_pos=grid&pfm_type=search_page&offerId=607862e60c070442666ca14f&buyboxToken=smartbuybox-acom-v2-addaba6c-8376-4596-8176-344fc36d2414-2021-10-12%2011%3A01%3A00-0300&cor=BRANCO&tamanho=17"
	americanasCrawler := crawlers.Americanas{}
	offer := americanasCrawler.GetOfferFor(url)

	fmt.Println(offer)
}
