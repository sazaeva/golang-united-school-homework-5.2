package cache

import (
	"time"
)

type Cache struct {
	value    string
	deadline time.Time
}

type MainCache struct {
	c map[string]Cache
}

func NewCache() MainCache {
	return MainCache{c: map[string]Cache{}}
}

func (c *MainCache) Put(key, value string) {

	c.c[key] = Cache{value: value, deadline: time.Now().AddDate(100, 0, 0)}
	return
}

func (c *MainCache) PutTill(key, value string, deadline time.Time) {

	c.c[key] = Cache{value: value, deadline: deadline}
	return
}

func (c *MainCache) Get(key string) (string, bool) {

	if c.c[key].deadline.Before(time.Now()) {
		return "", false
	}

	if val, ok := c.c[key]; ok {
		return val.value, true
	}

	return "", true
}

func (c *MainCache) Keys() []string {
	keys := []string{}
	for k, v := range c.c {
		if v.deadline.Before(time.Now()) {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}
