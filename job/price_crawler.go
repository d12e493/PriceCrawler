package job

import (
	momo "product-query/crawler/momo"
)

type PriceCrawlerJob struct {
}

func CreatePriceCrawlerJob() *PriceCrawlerJob {
	return &PriceCrawlerJob{}
}

func (self *PriceCrawlerJob) BeforeProcess() {
}

func (self *PriceCrawlerJob) Process(args []string) {

	momo.PageProcess()
	// pchome.PageProcess()
}

func (self *PriceCrawlerJob) AfterProcess() {
}
