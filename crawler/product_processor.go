package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	bo "product-query/bo"
	. "product-query/global"
)

//pass generic product info to MQ/DB/api
func SendProductInfo(product bo.CrawlerProductBO) {
	Logger.Debug("SendProductInfo")
	Logger.Debug(product)

	url := fmt.Sprintf("http://%s:%d/v1/product", Config.Api.Host, Config.Api.Port)

	productValue, _ := json.Marshal(product)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(productValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	Logger.Debug("response Status:", resp.Status)
	Logger.Debug("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	Logger.Debug("response Body:", string(body))
}
