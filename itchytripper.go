package main

import (
	"net/http"
	"net/url"
)

type ItchyTripper struct {
	url *url.URL
}

func (t ItchyTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	originReq, _ := url.Parse(req.RequestURI)
	url := t.url
	url.Path = originReq.Path
	req.RequestURI = ""
	req.URL = t.url
	req.Host = t.url.Host
	return client.Do(req)
}
