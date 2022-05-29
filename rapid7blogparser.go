package main

import (
	"log"
)

var filename = "news.json"

func main() {
	// Parsing Rapid7 Blog
	rapid_blog := make(map[string]string)
	rapid_blog = CollectRapidBlog()
	if rapid_blog == nil {
		log.Println("The Rapid Parse went wrong.")
	}
	// Parsing TheHackerNews
	hacker_news := make(map[string]string)
	hacker_news = CollectHackerNews()
	if hacker_news == nil {
		log.Println("The Hacker News Parse went wrong.")
	}
	//Merging our maps to compare with our JSON
	input_res := mapMerge(rapid_blog, hacker_news)
	if input_res == nil {
		log.Println("Can not merge two maps with news.")
	}
	// Reading JSON file with existing blogs
	file_res, err := mapReader(filename)
	if err != nil || file_res == nil {
		ferr := mapWriter(input_res, filename)
		if ferr != nil {
			log.Println("Can not Read or Write a new map for the file " + filename)
		}
	}
	// Comparing our blog and file inputs
	compare_maps, isequal := blogEqual(input_res, file_res)
	if !isequal || compare_maps == nil {
		log.Println("We have a new blog!")
		// If we've found a new blog we are merging new blogs with existing ones
		maptofile := mapMerge(compare_maps, file_res)
		// Marshaling and writing to the file
		err := mapWriter(maptofile, filename)
		if err != nil {
			log.Println("Can not Marshal or Write to the file " + filename)
		}
		// TODO: Make a telegram bot message for 'compare_maps'
	} else {
		log.Println("Nothing to add. Sleeping till the next time.")
	}
}
