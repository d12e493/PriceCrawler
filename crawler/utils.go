package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetResponse(url string) string {

	response, err := http.Get(url)
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
