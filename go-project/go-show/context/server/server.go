package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal((http.ListenAndServe("127.0.0.1:8080", nil)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Handler started")
	defer log.Printf("handler ended")

	select {
	case <-time.After(time.Second * 5):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
