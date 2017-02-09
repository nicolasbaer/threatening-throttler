package main

import (
	"testing"
	"time"
)

func TestThrottleThreatening(t *testing.T) {
	tt := NewThreteningThrottler(2, 5)

	tt.ThrottleThreatening(nil, "1")
	tt.ThrottleThreatening(nil, "2")
	tt.ThrottleThreatening(nil, "3")
	tt.ThrottleThreatening(nil, "4")
	tt.ThrottleThreatening(nil, "5")

	if ret, _ := tt.ThrottleThreatening(nil, "6"); ret {
		t.Error("should not be able to insert", ret)
	}
	if ret, _ := tt.ThrottleThreatening(nil, "2"); !ret {
		t.Error("should be able to insert")
	}

	time.Sleep(6 * time.Second)

	if ret, _ := tt.ThrottleThreatening(nil, "6"); !ret {
		t.Error("should be able to insert", ret)
	}

}
