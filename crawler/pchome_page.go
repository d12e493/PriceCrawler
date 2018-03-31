package crawler

import (
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type PchomePage struct {
}

func (page PchomePage) FindProduct(body io.ReadCloser) {
	fmt.Println("FindProduct")

	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	defer body.Close()

	// Find the review items
	doc.Find("dl").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
		fmt.Println(s.Html())
		fmt.Println(s.Find("dd.c2f").Html())

		// title := s.Find("dd.c2f").Find("a").Text()
		// price := s.Find("dd.c3f").Find("span.value").Text()
		// fmt.Printf("Product title: %s price: %s  \n", title, price)
	})
}
