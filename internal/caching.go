package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	kv map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	initCache := make(map[string]cacheEntry)
	cacheNew := Cache{
		kv: initCache,
		mu: sync.Mutex{},
	}
	cacheNew.reapLoop(interval)
	return &cacheNew
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	temp := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.kv[key] = temp
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	res, ok := c.kv[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return res.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			<-ticker.C
			c.mu.Lock()
			for key, entry := range c.kv {
				if time.Since(entry.createdAt) > interval {
					delete(c.kv, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
