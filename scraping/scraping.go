package scraping

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	res, err := http.Get("https://www.youtube.com/watch?v=2gsXTVkdndc")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		fmt.Println("Name: ", name)
	})
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		desc, _ := s.Find("meta[itemprop=\"description\"]").Attr("content")
		fmt.Println("Description: ", desc)
	})
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		date, _ := s.Find("meta[itemprop=\"datePublished\"]").Attr("content")
		fmt.Println("Date: ", date)
	})
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		views, _ := s.Find("meta[itemprop=\"interactionCount\"]").Attr("content")
		fmt.Println("Views: ", views)
	})
}
