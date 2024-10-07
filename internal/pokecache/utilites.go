package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()

	defer c.mu.Unlock()

	if _, exists := c.cache[key]; exists {
		return fmt.Errorf("Entry already exists")

	}

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	return nil

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.Lock()

	defer c.mu.Unlock()

	if entry, exists := c.cache[key]; exists {
		return entry.val, exists
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}

}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}
