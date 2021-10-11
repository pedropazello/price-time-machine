package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	url := "https://www.americanas.com.br/produto/148004923"

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

	doc, err := htmlquery.Parse(strings.NewReader(string(body)))

	if err != nil {
		panic(err)
	}

	a := htmlquery.FindOne(doc, "//div[1]/div/div/main/div[3]/div[2]/div[1]/div[1]/div")

	fmt.Printf(htmlquery.InnerText(a))
}
