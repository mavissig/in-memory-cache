package cache

import "time"

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	cache := &Cache{
		items:             make(map[string]item),
		cleanupInterval:   cleanupInterval,
		defaultExpiration: defaultExpiration,
	}

	if cleanupInterval > 0 {
		go cache.runGC()
	}

	return cache
}

func (c *Cache) Set(k string, v interface{}, duration time.Duration) {
	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	} else if duration > 0 {
		expiration = time.Now().Add(duration).Unix()
	}

	c.Lock()
	c.items[k] = item{
		value:      v,
		expiration: expiration,
	}
	c.Unlock()
}

func (c *Cache) Remove(k string) {
	c.Lock()
	delete(c.items, k)
	c.Unlock()
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	c.RLock()
	defer c.RUnlock()

	item, ok := c.items[k]
	if !ok {
		return nil, false
	}
	if item.expiration > 0 {
		if time.Now().UnixNano() > item.expiration {
			return nil, false
		}
	}

	return item.value, true
}

func (c *Cache) runGC() {
	for {
		<-time.After(c.cleanupInterval)

		c.gc()
	}
}

func (c *Cache) gc() {
	if c.items == nil {
		return
	}

	var cleanedPool []string

	c.RLock()
	for k, v := range c.items {
		if time.Now().UnixNano() > v.expiration {
			cleanedPool = append(cleanedPool, k)
		}
	}
	c.RUnlock()

	if len(cleanedPool) > 0 {
		c.Lock()
		for _, k := range cleanedPool {
			delete(c.items, k)
		}
		c.Unlock()
	}
}
