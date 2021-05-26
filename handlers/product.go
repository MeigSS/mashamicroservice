package handlers

import (
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
		p.productGET(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.productPOST(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) productPOST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST successfully")
	pl := &data.Product{}
	err := pl.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
	}
	p.l.Printf("pl: %#v", pl)
	data.AddProduct(pl)
}

func (p *Product) productGET(rw http.ResponseWriter, r *http.Request) {

	/* marshal
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
	*/

	// encode
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
