package pchome

import (
	"fmt"
	"net/http"
	. "product-query/global"
	"product-query/service"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var momoMainPage = "https://www.momoshop.com.tw/main/Main.jsp"

func PageProcess() {

	Logger.Debug("Momo page process start")

	// get Lgrp category url
	LgrpCategory := make([]string, 5)

	LgrpCategories := getDocumentFromUrl(momoMainPage, "div.subMenu ul li a")

	Logger.Debug(LgrpCategories)
	Logger.Debug("LgrpCategory : " + strings.Join(LgrpCategory, " "))

	// get Dgrp category url
	DgrpCategory := make([]string, 5)

	Logger.Debug("DgrpCategory : " + strings.Join(DgrpCategory, " "))

	// get product
	productWorkerService := service.CreateWorkerService(10, nil)

	productWorkerService.Wait()

	Logger.Debug("Momo page process finish")
}

func productWorker(value interface{}) {
	productlUrl := value.(string)
	Logger.Debug("productWorker " + productlUrl)
}

func getDocumentFromUrl(url string, selector string) []*goquery.Selection {
	res, err := http.Get(url)
	if err != nil {
		Logger.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		Logger.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	Logger.Debug("url : " + url)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		Logger.Error(err)
	}
	selectors := make([]*goquery.Selection, 5)

	Logger.Debug("selector : " + selector)

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		fmt.Println(s)
		selectors = append(selectors, s)
	})

	return selectors
}
