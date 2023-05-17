package main

import (
	w "github.com/Username/Project-Name/web/api"
	s "github.com/Username/Project-Name/web/scraping"
)

func main() {
	channel := "https://www.youtube.com/watch?v=LPpqcvIlrhQ"
	s.Scrape(channel)
	w.StartServer()
}
