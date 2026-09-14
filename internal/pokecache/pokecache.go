package pokecache

import (
	"sync"
	"time"
)

type cacheEntries struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]cacheEntries
	mu         sync.Mutex
	duration   time.Duration
}

func NewCache(interval time.Duration) *Cache {
	var newCache Cache
	c := &newCache
	newCache.cacheEntry = make(map[string]cacheEntries, 0)
	newCache.duration = interval
	go func() {
		c.reapLoop()
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	ce := cacheEntries{}
	ce.val = val
	ce.createdAt = time.Now()
	c.mu.Lock()
	c.cacheEntry[key] = ce
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cacheEntry[key]
	c.mu.Unlock()
	if ok {
		return entry.val, true
	}
	return entry.val, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for now := range ticker.C {
		c.mu.Lock()
		for k, v := range c.cacheEntry {
			if now.Sub(v.createdAt) > c.duration {
				delete(c.cacheEntry, k)
			}
		}
		c.mu.Unlock()
	}
}
