package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	url := "https://www.americanas.com.br/produto/2779138678?pfm_carac=bicicleta-caloi-vulcan&pfm_page=search&pfm_pos=grid&pfm_type=search_page&offerId=607862e60c070442666ca14f&buyboxToken=smartbuybox-acom-v2-addaba6c-8376-4596-8176-344fc36d2414-2021-10-12%2011%3A01%3A00-0300&cor=BRANCO&tamanho=17"

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

	bodyAsString := string(body)

	// Write File for debugging purpose
	arquivo, err := os.OpenFile("americanas.html", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(bodyAsString)
	arquivo.Close()
	//

	doc, err := htmlquery.Parse(strings.NewReader(bodyAsString))

	if err != nil {
		panic(err)
	}

	a, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/main/div[2]/div[2]/div[1]/div[2]/div")

	if err != nil {
		fmt.Printf("Price not found for URL:", url)
	}

	fmt.Printf(htmlquery.InnerText((a)))
}
