package main

import (
	"github.com/gocolly/colly"
	"log"
)

func CollectHackerNews() map[string]string {
	url := "https://thehackernews.com"
	domain := "thehackernews.com"
	telegram_msg := make(map[string]string)

	c := colly.NewCollector(colly.AllowedDomains(domain))
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(".body-post", func(e *colly.HTMLElement) {
		link := e.ChildAttrs("a", "href")
		subtitle := e.ChildText(".home-title")
		telegram_msg[subtitle] = link[0]
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)
	return telegram_msg
}
