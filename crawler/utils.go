package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
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
func Decodebig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
