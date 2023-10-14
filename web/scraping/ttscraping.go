package scraping

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	NameAcc, LikesAcc, SubsAcc, DescAcc string
)

func ScrapeAccount(accountURL string) {
	res, err := http.Get(accountURL)
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
	doc.Find("strong[data-e2e=\"followers-count\"]").Each(func(i int, s *goquery.Selection) {
		SubsAcc = s.Text()
	})
	doc.Find("strong[data-e2e=\"likes-count\"]").Each(func(i int, s *goquery.Selection) {
		LikesAcc = s.Text()
	})
	doc.Find("h2[data-e2e=\"user-bio\"]").Each(func(i int, s *goquery.Selection) {
		DescAcc = s.Text()
	})
	doc.Find("h1[data-e2e=\"user-title\"]").Each(func(i int, s *goquery.Selection) {
		NameAcc = s.Text()
	})
}
