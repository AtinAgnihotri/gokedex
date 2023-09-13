package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ICache interface {
	Add(key string, val CacheEntry)
	Get(key string) (val []byte, err error)
	reapLoop()
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	entries := make(map[string]CacheEntry)
	cache := Cache{
		entries: entries,
		mu:      &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		fmt.Println("Running REAP LOOP")
		for key, value := range c.entries {
			diff := time.Now().Sub(value.createdAt)

			if diff > interval {
				fmt.Println("Removing", key)
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
	}
}

func (c *Cache) Get(key string) (val []byte, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var empty []byte
	entry, ok := c.entries[key]
	if !ok {
		return empty, errors.New(fmt.Sprintf("No value found corresponding to key %v", key))
	}
	return entry.val, nil
}
