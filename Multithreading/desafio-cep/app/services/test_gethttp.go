package services

import (
	"testing"
	"time"
)

func TestGetCEPInfo(t *testing.T) {
	ch := make(chan interface{})
	go GetCEPInfo(12345678, "viacep", ch)
	go GetCEPInfo(12345678, "brasilapi", ch)

	select {
	case result := <-ch:
		if result == nil {
			t.Error("Expected a result, got nil")
		}
	case <-time.After(1 * time.Second):
		t.Error("Test timed out")
	}
}
