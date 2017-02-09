// Threatening-Throttler implements a throtteling reverse proxy
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "0.0.1"

var (
	verbose   bool   // -verbose flag
	localPort string // -port flag
)

func init() {
	flag.BoolVar(&verbose, "verbose", true, "verbose mode")
	flag.StringVar(&localPort, "port", ":8080", "local port accepting request")
}

func throttelHandler(rw http.ResponseWriter, req *http.Request) {
	http.Error(rw, "Sorry, still wip", http.StatusNotImplemented)
}

func main() {
	flag.Parse()
	if verbose {
		fmt.Println("This is Threatening-Throttler version", version)
	}

	http.HandleFunc("/", http.HandlerFunc(throttelHandler))
	log.Fatal(http.ListenAndServe(localPort, nil))

}
