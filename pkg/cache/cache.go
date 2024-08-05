package cache

func New() *Cache {
	return &Cache{
		Items: make(map[string]Item),
	}
}

func (c *Cache) Set(k string, v interface{}) {
	c.Lock()
	c.Items[k] = Item{
		value: v,
	}
	c.Unlock()
}

func (c *Cache) Remove(k string) {
	c.Lock()
	delete(c.Items, k)
	c.Unlock()
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	if v, ok := c.Items[k]; ok {
		return v.value, true
	}

	return nil, false
}
