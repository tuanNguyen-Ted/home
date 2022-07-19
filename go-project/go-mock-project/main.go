package main

import (
	crawler "go_mock_project/package/go_crawler"
	goS3 "go_mock_project/package/s3"
	"log"
	"time"
)

func main() {
	start := time.Now()
	//Enter URL here
	url := "https://fptshop.com.vn"

	//CODE here
	pageRecipe := crawler.Crawler(url)

	goS3.S3PutObject(url, pageRecipe)

	elapsed := time.Since(start)
	log.Printf("Exec took %s", elapsed)
}
