package cache

import (
	"sync"
	"time"
)

type item struct {
	created    time.Time
	expiration int64
	value      interface{}
}

type Cache struct {
	sync.RWMutex
	items             map[string]item
	cleanupInterval   time.Duration
	defaultExpiration time.Duration
}
