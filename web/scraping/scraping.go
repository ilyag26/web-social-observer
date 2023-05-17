package scraping

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var NameRet string
var DescRet string
var DateRet string
var ViewsRet string

func Scrape(Channel string) {
	res, err := http.Get(Channel)
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
		NameRet = s.Text()
	})
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		DescRet, _ = s.Find("meta[itemprop=\"description\"]").Attr("content")
		DateRet, _ = s.Find("meta[itemprop=\"datePublished\"]").Attr("content")
		ViewsRet, _ = s.Find("meta[itemprop=\"interactionCount\"]").Attr("content")
	})
}
