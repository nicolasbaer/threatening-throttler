package main

import (
	"net/http"
	"net/url"
)

type ItchyTripper struct {
	url    *url.URL
	client *http.Client
}

func (t ItchyTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	originReq, _ := url.Parse(req.RequestURI)
	url := t.url
	url.Path = originReq.Path
	req.RequestURI = ""
	req.URL = t.url
	req.Host = t.url.Host
	req.Close = true
	return t.client.Do(req)
}
