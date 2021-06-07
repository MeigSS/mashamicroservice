package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/masha/WebServer/data"
)

func (p *Products) ProductDELETE(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle DELETE")
	vars := mux.Vars(r)
	temp := vars["id"]
	id, err := strconv.Atoi(temp)
	if err != nil {
		http.Error(rw, "Analysing id failed!", http.StatusBadRequest)
	}
	data.DelProduct(id)
}
