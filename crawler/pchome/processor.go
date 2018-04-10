package pchome

import (
	"encoding/json"
	"fmt"
	bo "product-query/bo"
	"product-query/crawler"
	. "product-query/global"
	"product-query/service"
	"regexp"
	"strconv"
	"strings"
)

type PchomeMenu struct {
	Id string `json:"Id"`
}

type PchomeSunMenu struct {
	Id    string          `json:"Id"`
	Name  string          `json:"Name"`
	Nodes []PchomeSunMenu `json:"Nodes"`
}

type PchomeProduct struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Url  string
	Pic  struct {
		B string `json:"B"`
		S string `json:"S"`
	} `json:"Pic"`
	Price struct {
		M int `json:"M"`
		P int `json:"P"`
	} `json:"Price"`
}

var rgx = regexp.MustCompile(`count\(\d\)`)
var productUrl string = "https://24h.pchome.com.tw/prod/"
var productPictureUrl string = "https://a.ecimg.tw"
var productCountUrl string = "http://ecapi.pchome.com.tw/ecshop/prodapi/v2/store/${subMenu}/prod/count&_callback=jsonp_prodcount"
var productDetailUrl string = "http://ecapi.pchome.com.tw/ecshop/prodapi/v2/store/${subMenu}/prod&offset=1&limit=${count}&fields=Id,Nick,Pic,Price,Discount,isSpec,Name,isCarrier,isSnapUp,isBigCart&_callback=jsonp_prodlist"

func PageProcess() {

	Logger.Debug("Pchome page process start")

	categorys := []string{"daily", "food", "life", "outdoor", "beauty", "vogue", "3c", "cp", "nb", "mobile", "digi", "ce"}

	var menuUrl string = "http://ecapi.pchome.com.tw/cdn/ecshop/cateapi/v1.5/region&sign=h24%252F${category}&_callback=cb_ecshopCategoryRegion&25374809"
	menus := make([]PchomeMenu, 5)

	// // get menu
	for _, category := range categorys {
		url := strings.Replace(menuUrl, "${category}", category, -1)
		tmpMenus := make([]PchomeMenu, 5)
		jsonParseFromUrl(url, &tmpMenus)
		menus = append(menus, tmpMenus...)
	}
	// get submenu
	var subMenuUrl string = "http://ecapi.pchome.com.tw/cdn/ecshop/cateapi/v1.5/region/${menu}/menu&_callback=jsonp_nemu&25374802?_callback=jsonp_nemu"
	subMenus := make([]PchomeSunMenu, 5)

	for _, menu := range menus {
		if len(menu.Id) > 0 {
			url := strings.Replace(subMenuUrl, "${menu}", menu.Id, -1)
			tmpMenus := make([]PchomeSunMenu, 5)
			jsonParseFromUrl(url, &tmpMenus)
			subMenus = append(subMenus, tmpMenus...)
		}
	}

	// get product
	productWorkerService := service.CreateWorkerService(10, productWorker)

	for _, subMenu := range subMenus {
		getProductCountAndDetail(subMenu.Id, productWorkerService)
		if len(subMenu.Nodes) > 0 {
			for _, childSubMenu := range subMenu.Nodes {
				getProductCountAndDetail(childSubMenu.Id, productWorkerService)
			}
		}
	}
	productWorkerService.Wait()

	Logger.Debug("Pchome page process finish")
}

// For Test
// PageProcess
func PageProcessTest() {
	productWorkerService := service.CreateWorkerService(10, productWorker)
	getProductCountAndDetail("DAAG66", productWorkerService)
	productWorkerService.Wait()
}

func getProductCountAndDetail(subMenuId string, workService *service.WorkerService) {
	var url string
	if len(subMenuId) > 0 {
		count := getProductCount(subMenuId)
		if count > 0 {
			r := strings.NewReplacer("${subMenu}", subMenuId,
				"${count}", strconv.Itoa(count))
			url = r.Replace(productDetailUrl)
			workService.Submit(url)
		}
	}
}
func getProductCount(subMenuId string) int {
	url := strings.Replace(productCountUrl, "${subMenu}", subMenuId, -1)
	response := crawler.GetResponse(url)

	rs := rgx.FindStringSubmatch(response)

	if len(rs) > 0 {
		count, err := strconv.Atoi(rs[0])

		if err != nil {
			fmt.Println(err)
			return 0
		}
		return count
	}

	return 100
}
func productWorker(value interface{}) {
	productlUrl := value.(string)
	products := make([]PchomeProduct, 5)

	jsonParseFromUrl(productlUrl, &products)

	for _, product := range products {
		p := bo.CrawlerProductBO{
			Name:          product.Name,
			Website:       "pchome",
			SiteProductId: product.Id,
			Link:          productUrl + product.Id,
			Price:         product.Price.P,
			Picture:       productPictureUrl + product.Pic.B,
		}
		crawler.SendProductInfo(p)
	}
}

func jsonParseFromUrl(url string, value interface{}) {
	response := crawler.GetResponse(url)
	responseArray := getArrayFromResponse(response)
	if len(responseArray) > 0 {
		Logger.Debug(responseArray)
		json.Unmarshal([]byte(responseArray), value)
	} else {
		Logger.Debug("no responseArray")
	}
}
func getArrayFromResponse(response string) string {
	start := strings.Index(response, "[")
	end := strings.LastIndex(response, "]")
	if start > 0 && end > 0 {
		return response[start : end+1]
	}
	return ""
}
