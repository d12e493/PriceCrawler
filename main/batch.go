package main

import (
	"log"
	"os"
	. "product-query/global"
	job "product-query/job"
)

var jobMap = map[string]Job{
	"priceCrawler": job.CreatePriceCrawlerJob(),
}

type Job interface {
	BeforeProcess()
	Process(args []string)
	AfterProcess()
}

var jobName string = Arguments.JobName

func init() {

	if logFile, err := CreateLogFile("batch"); err == nil {
		log.SetOutput(logFile)
	}

	if jobName == "" {
		panic("no jobname arguments")
	}
}

func main() {

	// get job
	var job Job = jobMap[jobName]

	if job != nil {
		log.Println("[Start Job] " + jobName)

		// process
		job.Process(os.Args)
		defer job.AfterProcess()

		log.Println("[Finish Job] " + jobName)
	} else {
		panic(" job '" + jobName + "' is not exist !!! ")
	}
}
