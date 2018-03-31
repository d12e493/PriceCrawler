package pchome

import (
	"encoding/json"
	"fmt"
	"product-query/crawler"
	"strings"
)

type PchomeMenu struct {
	Id string `json:"Id"`
}

func PageProcess() {
	categorys := []string{"daily", "food"}

	var menuUrl string = "http://ecapi.pchome.com.tw/cdn/ecshop/cateapi/v1.5/region&sign=h24%252F${category}&_callback=cb_ecshopCategoryRegion&25374809"

	var menuSlice []string
	for _, category := range categorys {
		url := strings.Replace(menuUrl, "${category}", category, -1)
		fmt.Println(url)
		response := crawler.GetResponse(url)
		responseArray := getArrayFromResponse(response)
		menus := make([]PchomeMenu, 5)
		json.Unmarshal([]byte(responseArray), &menus)

		if len(menus) > 0 {
			for _, m := range menus {
				menuSlice = append(menuSlice, m.Id)
			}
		}
	}
	fmt.Println(menuSlice)
	// var subMenuUrl string = "http://ecapi.pchome.com.tw/cdn/ecshop/cateapi/v1.5/region/${menu}/menu&_callback=jsonp_nemu&25374802?_callback=jsonp_nemu"
}

func getArrayFromResponse(response string) string {
	start := strings.Index(response, "[")
	end := strings.Index(response, "]")
	return response[start : end+1]
}
