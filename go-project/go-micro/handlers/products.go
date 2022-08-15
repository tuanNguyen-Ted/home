package handlers

import (
	"log"
	"my-micro/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.GetProducts(w, r)
	case http.MethodPost:
		p.CreateProducts(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusInternalServerError)
	}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) CreateProducts(w http.ResponseWriter, r *http.Request) {

}
