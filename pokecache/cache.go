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
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entry, key)
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for tick := range ticker.C {
		c.reap(time.Duration(tick.Second()))
	}
}

func (c *Cache) reap(interval time.Duration) {
	for k, v := range c.entry {
		if time.Since(v.createdAt).Seconds() > interval.Seconds() {
			c.delete(k)
		}
	}
}
