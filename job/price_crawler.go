package job

import (
	momo "product-query/crawler/momo"
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

	// momo.PageProcess()
	momo.PageProcessTest()
	// pchome.PageProcess()
	pchome.PageProcessTest()
}

func (self *PriceCrawlerJob) AfterProcess() {
}
