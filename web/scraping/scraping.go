package scraping

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	d "github.com/Username/Project-Name/web/data"
)

var (
	NameVideoRet, ViewsVideoRet, DescVideoRet, DateVideoRet, NameChannelRet, ViewsChannelRet, DescChannelRet, DateChannelRet, SubsChannelRet string
)
var ArrayForData []string
var Views string

func ScrapeVideo(videoURL string) {
	res, err := http.Get(videoURL)
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
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		NameVideoRet, _ = s.Find("meta[itemprop=\"name\"]").Attr("content")
		ViewsVideoRet, _ = s.Find("meta[itemprop=\"description\"]").Attr("content")
		DescVideoRet, _ = s.Find("meta[itemprop=\"datePublished\"]").Attr("content")
		DateVideoRet, _ = s.Find("meta[itemprop=\"interactionCount\"]").Attr("content")
	})
}

func ScrapeChannel(channelURL string) {
	ArrayForData = nil

	res, err := http.Get(channelURL + "/about")
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
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		NameChannelRet, _ = s.Find("meta[property=\"og:title\"]").Attr("content")
		DescChannelRet, _ = s.Find("meta[property=\"og:description\"]").Attr("content")
	})
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		ArrayForData = append(ArrayForData, s.Text())
	})
	dataValue := strings.Trim(ArrayForData[34], "var ytInitialData = ;")
	channelData := d.AutoGenerated{}
	json.Unmarshal([]byte(dataValue), &channelData)
	for _, row := range channelData.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		for _, row2 := range row.TabRenderer.Content.SectionListRenderer.Contents {
			for _, row3 := range row2.ItemSectionRenderer.Contents {
				ViewsChannelRet = row3.ChannelAboutFullMetadataRenderer.ViewCountText.SimpleText
				for _, row4 := range row3.ChannelAboutFullMetadataRenderer.JoinedDateText.Runs {
					DateChannelRet = row4.Text
				}
			}
		}
	}
	SubsChannelRet = channelData.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText
}
