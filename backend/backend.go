package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\u0950 goa \u0950", "test")
	r.Body.Close()
}

func main() {
	backend := &http.Server{
		Addr:         ":8081",
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	backend.SetKeepAlivesEnabled(false)
	backend.ListenAndServe()

}
