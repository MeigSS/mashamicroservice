package handlers

import (
	"net/http"

	"github.com/masha/WebServer/data"
)

// swagger:route POST /products products postProduct
// Return None

func (p *Products) ProductPOST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Just a test, post work well")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
