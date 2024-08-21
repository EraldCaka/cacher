package cacher

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string][]byte
}

type Cacher interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte, expiration time.Duration) error
	Has(key []byte) bool
	Delete(key []byte) error
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	strKey := string(key)
	value, ok := c.data[strKey]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", strKey)
	}
	return value, nil
}

func (c *Cache) Set(key []byte, value []byte, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	strKey := string(key)
	c.data[strKey] = value

	if expiration > 0 {
		go func() {
			<-time.After(expiration)
			c.mu.Lock()
			defer c.mu.Unlock()
			delete(c.data, strKey)
		}()
	}
	return nil
}

func (c *Cache) Has(key []byte) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.data[string(key)]

	return ok
}

func (c *Cache) Delete(key []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, string(key))
	return nil
}
