package cache

import "sync"

type Item struct {
	value interface{}
}

type Cache struct {
	sync.RWMutex
	Items map[string]Item
}
