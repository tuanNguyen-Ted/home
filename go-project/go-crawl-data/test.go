package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/temoto/robotstxt"
)

func main() {
	resp, err := http.Get("https://tiki.vn/robots.txt")
	robots, err := robotstxt.FromResponse(resp)
	resp.Body.Close()
	if err != nil {
		log.Println("Error parsing robots.txt:", err.Error())
	}

	fmt.Println("robot:", robots.Sitemaps)
}
