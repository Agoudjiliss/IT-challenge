package server

import (
    "sync"
    "time"
)

// Cache structure to store cached responses with expiration
type Cache struct {
    Data     map[string]*CacheItem
    ExpireIn time.Duration
    Lock     sync.RWMutex
}

// Cache item with expiration time
type CacheItem struct {
    Response  []byte
    ExpiredAt time.Time
}

// Create a new cache
func NewCache(expireIn time.Duration) *Cache {
    return &Cache{
        Data:     make(map[string]*CacheItem),
        ExpireIn: expireIn,
    }
}

// Get the cached item
func (c *Cache) Get(key string) ([]byte, bool) {
    c.Lock.RLock()
    defer c.Lock.RUnlock()

    item, exists := c.Data[key]
    if !exists || time.Now().After(item.ExpiredAt) {
        return nil, false
    }

    return item.Response, true
}

// Set the cache item with expiration
func (c *Cache) Set(key string, response []byte) {
    c.Lock.Lock()
    defer c.Lock.Unlock()

    c.Data[key] = &CacheItem{
        Response:  response,
        ExpiredAt: time.Now().Add(c.ExpireIn),
    }
}
