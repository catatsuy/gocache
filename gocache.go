package gocache

import "sync"

type cache struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() *cache {
	m := make(map[string]interface{})
	c := &cache{
		items: m,
	}
	return c
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.RLock()
	v, found := c.items[key]
	c.RUnlock()
	return v, found
}

func (c *cache) Set(key string, value interface{}) {
	c.Lock()
	c.items[key] = value
	c.Unlock()
}

func (c *cache) Incr(key string, n int) {
	c.Lock()
	v, found := c.items[key]
	if found {
		c.items[key] = v.(int) + n
	} else {
		c.items[key] = n
	}
	c.Unlock()
}
