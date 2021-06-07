package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/masha/WebServer/handlers"
)

func main() {
	// New a Logger with description @param1 and tag @param2
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// New a product, which implement http request function
	product_handle := handlers.NewProduct(l)

	// New a router instance
	sm := mux.NewRouter()

	// Methods register a new [route] with a matcher for http methods.
	// Subrouter create a [subrouter] for the route.
	// Subrouter will test inner routes only if the parent route matched.
	//						this is a router
	//							|------- "/" goes here
	//							|------- "/ppp" goes here if registered
	// "GET"(this is a route) ---------- .............
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", product_handle.ProductGET)

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/{id:[0-9]+}", product_handle.ProductDELETE)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", product_handle.ProductPOST)
	postRouter.Use(product_handle.MiddlewareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", product_handle.ProductPUT)
	putRouter.Use(product_handle.MiddlewareValidateProduct)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// listen and catch signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	// Cancel context and release resources
	timeout_context, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeout_context)

	// server mux also implement the handler interface
	// http.ListenAndServe(":9090", sm)
}
