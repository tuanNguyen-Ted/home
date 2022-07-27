package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	// defer cancel()

	// req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req = req.WithContext(ctx)
	// res, err := http.DefaultClient.Do(req)
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
