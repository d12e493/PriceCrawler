package job

import (
	"product-query/crawler/pchome"
)

type PriceCrawlerJob struct {
}

func CreatePriceCrawlerJob() *PriceCrawlerJob {
	return &PriceCrawlerJob{}
}

func (self *PriceCrawlerJob) BeforeProcess() {
}

func (self *PriceCrawlerJob) Process(args []string) {

	pchome.PageProcess()
}

func (self *PriceCrawlerJob) AfterProcess() {
}
