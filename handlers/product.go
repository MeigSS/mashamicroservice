package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masha/WebServer/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ProductGET(rw http.ResponseWriter, r *http.Request) {
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

func (p *Products) ProductPOST(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Just a test, post work well")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) ProductPUT(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["id"]
	id, _ := strconv.Atoi(temp)

	p.l.Println("Handle PUT")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.UpdateProduct(id, &prod)
	fmt.Println("id is ", id)
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
