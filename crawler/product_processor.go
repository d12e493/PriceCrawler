package crawler

import (
	"fmt"
)

type CrawlerProduct struct {
	Name          string
	Website       string
	SiteProductId string
	Link          string
	Price         int
	Picture       string
}

//pass generic product info to MQ/DB/api
func SendProductInfo(product CrawlerProduct) {
	fmt.Println("SendProductInfo")
	fmt.Println(product)
}
