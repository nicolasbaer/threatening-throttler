// Threatening-Throttler implements a throtteling reverse proxy
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const version = "0.0.1"
const cookieName = "threatening-cookie"

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

	cookie, err := req.Cookie(cookieName)
	var id string
	if err == nil {
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("Error :( ", err)
			return
		}

		id = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	} else {
		id = cookie.Value
	}

	allow, err := ThrottleThreatening(req, id)
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

	u, _ := url.Parse("http://nginx.org/")

	proxy = httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = ItchyTripper{u}

	http.HandleFunc("/", http.HandlerFunc(throttelHandler))
	log.Fatal(http.ListenAndServe(localPort, nil))

}
