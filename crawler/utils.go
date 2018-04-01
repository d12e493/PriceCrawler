package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResponse(url string) string {

	time.Sleep(1 * time.Second)

	response, err := http.Get(url)

	fmt.Println(url)
	fmt.Println(response.StatusCode)

	if response.StatusCode != 200 {
		time.Sleep(5 * time.Second)
		return GetResponse(url)
	}

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		return string(contents)
	}

	return ""
}
