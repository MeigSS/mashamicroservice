package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/masha/WebServer/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Handle an GET
	if r.Method == http.MethodGet {
		p.convertJSONencode(rw, r)
		return
	}
	// Handle an update
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// using marshal, call in ServeHTTP
func (p *Product) convertJSONmarshal(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}

// using encode, call in ServeHTTP
func (p *Product) convertJSONencode(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
