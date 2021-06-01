package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/masha/WebServer/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.productGET(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.productPOST(rw, r)
		return
	}
	if r.Method == http.MethodPut {

		p.l.Println("Handle PUT")
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI1")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI2")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI3")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("got id ", id)

		p.productPUT(id, rw, r)

		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) productGET(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	/*// using mashal
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
	*/

	// using encode
	// this will be faster than using mashal
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) productPOST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Just a test, post work well")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable unmarshal JSON", http.StatusBadRequest)
		return
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) productPUT(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	fmt.Printf("%#v", prod)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
		return
	}

	data.UpdateProduct(id, prod)
}
