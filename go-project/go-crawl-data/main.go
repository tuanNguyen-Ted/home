package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/temoto/robotstxt"
)

var (
	url       = "https://tiki.vn/"
	maxDepth  = 1
	knownUrls = []string{}
)

func main() {
	c := colly.NewCollector(colly.MaxDepth(maxDepth))

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	//Search for robots.txt
	url = strings.TrimSpace(url)
	var robotUrl string
	if url[len(url)-1] != '/' {
		robotUrl = url + "/robots.txt"
	} else {
		robotUrl = url + "robots.txt"
	}

	fmt.Printf("robots.txt URL: %v\n", robotUrl)
	resp, err := http.Get(robotUrl)
	robots, err := robotstxt.FromResponse(resp)
	resp.Body.Close()

	if err != nil {
		log.Println("Error parsing robots.txt:", err.Error())
	}

	fmt.Println("robots.Sitemaps :", robots.Sitemaps)
	fmt.Println("amounts robots.Sitemaps :", len(robots.Sitemaps))
	if len(robots.Sitemaps) <= 0 {
		log.Println("No sitemaps found!")
		return
	}

	for _, sitemapUrl := range robots.Sitemaps {
		c.Visit(sitemapUrl)
	}

	PrintLog(knownUrls)
}

func PrintLog(knownUrls []string) {
	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
	}
	fmt.Println("Collected", len(knownUrls), "URLs")
}
