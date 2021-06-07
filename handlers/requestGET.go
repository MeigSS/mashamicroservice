package handlers

import (
	"net/http"

	"github.com/masha/WebServer/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponse

func (p *Products) ProductGET(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
