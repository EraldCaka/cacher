package cacher

import (
	"testing"
	"time"
)

func TestCacheIntegration(t *testing.T) {
	cache := New()

	key := []byte("testKey")
	value := []byte("testValue")
	err := cache.Set(key, value, 0)
	if err != nil {
		t.Errorf("Unexpected error during Set: %v", err)
	}

	retrievedValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Unexpected error during Get: %v", err)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %s, but got %s", value, retrievedValue)
	}

	ttl := time.Millisecond * 100
	err = cache.Set(key, value, ttl)
	if err != nil {
		t.Errorf("Unexpected error during Set with TTL: %v", err)
	}

	time.Sleep(ttl + time.Millisecond*50)

	if cache.Has(key) {
		t.Error("Expected key to be expired, but it's still present")
	}

	err = cache.Delete(key)
	if err != nil {
		t.Errorf("Unexpected error during Delete: %v", err)
	}

	if cache.Has(key) {
		t.Error("Expected key to be deleted, but it's still present")
	}
}
