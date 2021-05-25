package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/masha/WebServer/handlers"
)

func main() {

	l := log.New(os.Stdout, "test-api", log.LstdFlags)
	hello_handle := handlers.NewHello(l)
	product_handle := handlers.NewProduct(l)
	goodbye_handle := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hello_handle)
	sm.Handle("/goodbye", goodbye_handle)
	sm.Handle("/product", product_handle)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// this will block
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	timeout_context, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeout_context)
}
