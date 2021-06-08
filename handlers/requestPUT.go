package handlers

import (
	"net/http"

	"github.com/masha/WebServer/data"
)

// swagger:route PUT /products products putProduct
// Return None

func (p *Products) ProductPUT(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	id := prod.ID
	data.UpdateProduct(id, &prod)
}
