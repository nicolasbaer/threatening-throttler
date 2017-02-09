package main

import (
	"math/rand"
	"net/http"
)

func ThrottleThreatening(req *http.Request, id string) (bool, error) {
	return rand.Int()%2 == 0, nil
}
