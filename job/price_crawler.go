package job

import (
	"fmt"
	"product-query/bo"
	"product-query/crawler/pchome"
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

	pchome.PageProcess()
}

func (self *PriceCrawlerJob) AfterProcess() {
}
