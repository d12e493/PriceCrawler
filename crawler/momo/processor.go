package pchome

import (
	"net/http"
	"product-query/bo"
	"product-query/crawler"
	. "product-query/global"
	"product-query/service"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var momoMainPage = "http://www.momoshop.com.tw/main/Main.jsp"
var productUrl = "http://www.momoshop.com.tw"

func PageProcess() {

	Logger.Debug("Momo page process start")

	// get Lgrp category url
	LgrpCategoryUrlSlice := make([]string, 0)

	LgrpCategories := getDocumentFromUrl(momoMainPage, "div.subMenu li.BTDME a")
	for _, LgrpCategoy := range LgrpCategories {
		if LgrpCategoy != nil {
			href, _ := LgrpCategoy.Attr("href")
			if strings.Contains(href, "LgrpCategory.jsp") {
				LgrpCategoryUrlSlice = append(LgrpCategoryUrlSlice, href)
			}
		}
	}
	Logger.Debug("LgrpCategory : " + strings.Join(LgrpCategoryUrlSlice, " "))

	// get Dgrp category url
	DgrpCategoryUrlSlice := make([]string, 0)
	for _, Lurl := range LgrpCategoryUrlSlice {
		Logger.Debug("Lurl : " + Lurl)
		if len(Lurl) > 0 {
			DgrpCategories := getDocumentFromUrl(Lurl, "li[class^=cat] a")
			for _, DgrpCategory := range DgrpCategories {
				if DgrpCategory != nil {
					Logger.Debug("DgrpCategory")
					Logger.Debug(DgrpCategory)
					href, _ := DgrpCategory.Attr("href")
					if strings.Contains(href, "DgrpCategory.jsp") {
						DgrpCategoryUrlSlice = append(DgrpCategoryUrlSlice, productUrl+href)
					}
				}
			}
		}
	}

	Logger.Debug("DgrpCategory : " + strings.Join(DgrpCategoryUrlSlice, " "))

	// get product
	productWorkerService := service.CreateWorkerService(10, productWorker)

	for _, DgrpCategoryUrl := range DgrpCategoryUrlSlice {
		productWorkerService.Submit(DgrpCategoryUrl)
	}

	productWorkerService.Wait()

	Logger.Debug("Momo page process finish")
}

// for test
// func PageProcess() {
// 	productWorkerService := service.CreateWorkerService(10, productWorker)
// 	productWorkerService.Submit("https://www.momoshop.com.tw/category/DgrpCategory.jsp?d_code=2701800171&p_orderType=1")
// 	productWorkerService.Wait()
// }

func productWorker(value interface{}) {
	DgrpCategoryUrl := value.(string)
	products := getDocumentFromUrl(DgrpCategoryUrl, "li.eachGood")
	for _, product := range products {
		if product != nil {
			name, _ := product.ChildrenFiltered("a").Attr("title")
			id, _ := product.Find("input[name=goodsCode]").Attr("value")
			link, _ := product.ChildrenFiltered("a").Attr("href")
			priceStr := product.Find("span").Find("b").Text()
			price, _ := strconv.Atoi(priceStr)
			picture, _ := product.Find("img").Attr("src")

			nameC, _ := crawler.Decodebig5([]byte(name))
			name = string(nameC)
			p := bo.CrawlerProductBO{
				Name:          name,
				Website:       "momo",
				SiteProductId: id,
				Link:          productUrl + link,
				Price:         price,
				Picture:       "https" + picture,
			}
			Logger.Debug("productWorker")
			Logger.Debug(p)
			crawler.SendProductInfo(p)
		}
	}
}

func getDocumentFromUrl(url string, selector string) []*goquery.Selection {

	if !strings.Contains(url, "http") {
		url = "http:" + url
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(url)
	if err != nil {
		Logger.Error(err)
		return nil
	} else {

		defer res.Body.Close()
		if res.StatusCode != 200 {
			Logger.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		Logger.Debug("url : " + url)

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			Logger.Error(err)
			return nil
		} else {
			selectors := make([]*goquery.Selection, 5)

			Logger.Debug("selector : " + selector)

			doc.Find(selector).Each(func(i int, s *goquery.Selection) {
				selectors = append(selectors, s)
			})

			return selectors
		}
	}
}
