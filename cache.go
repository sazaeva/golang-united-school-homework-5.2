package cache

import (
	"sync"
	"time"
)

type Info struct {
	value    string
	deadline *time.Time
}

type Cache struct {
	mutex sync.Mutex
	c     map[string]Info
}

func NewCache() Cache {
	return Cache{c: map[string]Info{}}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.c[key]; ok {
		return "", false
	}
	return "", true
}

func (c *Cache) Put(key, value string) {
	c.mutex.Lock()
	c.c[key] = Info{value, nil}
	c.mutex.Unlock()
	return
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for k, _ := range c.c {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.c[key] = Info{value, &deadline}
}
