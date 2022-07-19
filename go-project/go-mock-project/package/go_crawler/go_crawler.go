package go_crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/temoto/robotstxt"
)

type PageRecipe struct {
	Url         string
	SitemapUrls []string
	Titles      []string
	ImgSrcs     []string
}

var (
	wg sync.WaitGroup
)

func Crawler(url string) (pageRecipe PageRecipe) {
	//Clean url
	c := colly.NewCollector()
	url = strings.TrimSpace(url)
	pageRecipe.Url = url

	defer func(PageRecipe) {
		fmt.Println("=====End of process=====")
		fmt.Printf("Page url crawled: %+v\n", pageRecipe.Url)
		fmt.Printf("Page titles crawled: %+v\n", pageRecipe.Titles)
		fmt.Printf("Page sitemapUrls crawled: %+v\n", len(pageRecipe.SitemapUrls))
		fmt.Printf("Page imgSrcs crawled: %+v\n", len(pageRecipe.ImgSrcs))
	}(pageRecipe)

	defer wg.Wait()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnXML("//loc", func(e *colly.XMLElement) {
		pageRecipe.SitemapUrls = append(pageRecipe.SitemapUrls, e.Text)
	})
	c.OnHTML("title, img", func(e *colly.HTMLElement) {
		switch e.Name {
		case "title":
			pageRecipe.Titles = append(pageRecipe.Titles, e.Text)
		case "img":
			pageRecipe.ImgSrcs = append(pageRecipe.ImgSrcs, e.Attr("src"))
			// Add more attribute per case
		}
	})
	c.OnScraped(func(r *colly.Response) {
		wg.Done()
		fmt.Println("Finished", r.Request.URL)
	})

	//Get sitemaps
	sitemapUrls := getSitemapUrls(url)
	for _, sitemapUrl := range sitemapUrls {
		wg.Add(1)
		go c.Visit(sitemapUrl)
	}

	//Get content page
	wg.Add(1)
	c.Visit(pageRecipe.Url)
	c.Clone()
	return
}

func getSitemapUrls(url string) []string {
	// Get url robots.txt
	var robotUrl string
	if url[len(url)-1] != '/' {
		robotUrl = url + "/robots.txt"
	} else {
		robotUrl = url + "robots.txt"
	}
	resp, err := http.Get(robotUrl)
	if err != nil {
		log.Println("robots.txt: URL", robotUrl)
		log.Println("Error when call get robots.txt: ", err.Error())
	}
	robots, err := robotstxt.FromResponse(resp)
	resp.Body.Close()
	if err != nil {
		log.Println("Error parsing robots.txt:", err.Error())
	}
	fmt.Printf("Robots.txt Sitemap URLs: %v\n", robots.Sitemaps)
	return robots.Sitemaps
}
