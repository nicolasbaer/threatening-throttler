// Threatening-Throttler implements a throtteling reverse proxy
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const version = "0.0.1"

var (
	verbose   bool   // -verbose flag
	localPort string // -port flag
)

var proxy *httputil.ReverseProxy

func init() {
	flag.BoolVar(&verbose, "verbose", true, "verbose mode")
	flag.StringVar(&localPort, "port", ":8080", "local port accepting request")
}

func throttelHandler(rw http.ResponseWriter, req *http.Request) {
	allow, err := ThrottleRandom(req)
	if err != nil {
		return
	}

	if !allow {
		fmt.Fprintln(rw, "<!DOCTYPE html><html><head></head><body><h1>access denied</h1></body></html>")
		return
	}

	// hand over to Reverse Proxy
	proxy.ServeHTTP(rw, req)

}

func main() {
	flag.Parse()
	if verbose {
		fmt.Println("This is Threatening-Throttler version", version)
	}

	u, _ := url.Parse("http://localhost:8081")

	proxy = httputil.NewSingleHostReverseProxy(u)

	http.HandleFunc("/", http.HandlerFunc(throttelHandler))
	log.Fatal(http.ListenAndServe(localPort, nil))

}
