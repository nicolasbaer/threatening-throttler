package main

import (
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type ThreteningThrottler struct {
	cache *cache.Cache
	max   int
}

func NewThreteningThrottler(max, ttl int) *ThreteningThrottler {
	t := ThreteningThrottler{
		cache: cache.New(time.Duration(ttl)*time.Second, 1*time.Second),
		max:   max,
	}

	return &t
}

func (t *ThreteningThrottler) ThrottleThreatening(req *http.Request, id string) (bool, error) {

	_, found := t.cache.Get(id)

	if found || len(t.cache.Items()) <= t.max {
		t.cache.SetDefault(id, nil)
		return true, nil
	} else {
		return false, nil
	}
}
