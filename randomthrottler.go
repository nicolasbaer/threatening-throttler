package main

import (
	"math/rand"
	"net/http"
)

func ThrottleRandom(req *http.Request) (bool, error) {
	return rand.Int()%2 == 0, nil
}
