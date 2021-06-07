package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masha/WebServer/data"
)

func (p *Products) ProductPUT(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp := vars["id"]
	id, _ := strconv.Atoi(temp)

	p.l.Println("Handle PUT")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.UpdateProduct(id, &prod)
	fmt.Println("id is ", id)
}
