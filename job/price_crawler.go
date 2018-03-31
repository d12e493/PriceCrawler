package job

import (
	"fmt"
	"net/http"
	"product-query/bo"
	"product-query/crawler"
	"product-query/service"
)

type PriceCrawlerJob struct {
}

func CreatePriceCrawlerJob() *PriceCrawlerJob {
	return &PriceCrawlerJob{}
}

func (self *PriceCrawlerJob) BeforeProcess() {
}

func (self *PriceCrawlerJob) ProductProcessor(value interface{}) {
	product := value.(bo.ProductBO)
	fmt.Println(product)
}

func (self *PriceCrawlerJob) Process(args []string) {
	productExecutor := service.CreateWorkerService(10, self.ProductProcessor)

	fmt.Println(productExecutor)

	pchomeUrlSlice := []string{"http://24h.pchome.com.tw/store/DSAA35?style=2"}
	tkecUrlSlice := []string{"https://www.tkec.com.tw/"}
	momoUrlSlice := []string{"www.momoshop.com.tw/index.html"}

	fmt.Println(pchomeUrlSlice, tkecUrlSlice, momoUrlSlice)

	for _, url := range pchomeUrlSlice {
		fmt.Println("crawler url : " + url)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		var pchome = crawler.PchomePage{}
		pchome.FindProduct(response.Body)
	}
}

func (self *PriceCrawlerJob) AfterProcess() {
}
