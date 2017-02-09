package main

import (
	"net/http"
	"testing"
)

func TestThrottleRandom(t *testing.T) {
	req := http.Request{}

	_, err := ThrottleRandom(&req)
	if err != nil {
		t.Error("we don't know what has happended")
	}
}
