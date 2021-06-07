package handlers

import (
	"net/http"

	"github.com/masha/WebServer/data"
)

func (p *Products) ProductPOST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Just a test, post work well")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
