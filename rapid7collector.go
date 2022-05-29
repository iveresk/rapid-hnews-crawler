package main

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func CollectRapidBlog() map[string]string {
	url := "https://www.rapid7.com/blog/posts"
	domain := "www.rapid7.com"
	telegram_msg := make(map[string]string)

	c := colly.NewCollector(colly.AllowedDomains(domain))
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		path := "https://" + domain + link
		if strings.Contains(link, "blog/post/") {
			subtitle := e.ChildText("h3")
			telegram_msg[subtitle] = path
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)
	return telegram_msg
}
