package cache

import (
	"context"
	"hash/maphash"
	"sync"

	"github.com/decimalbell/akv"
	"github.com/decimalbell/akv/internal"
)

const n = 512

type Cache struct {
	rwmutexs [n]*sync.RWMutex
	caches   [n]akv.Cache
}

func New() akv.Cache {
	c := new(Cache)
	for i := 0; i < n; i++ {
		c.rwmutexs[i] = new(sync.RWMutex)
		c.caches[i] = internal.NewCache()
	}
	return c
}

func (c *Cache) rwmutexcache(key string) (*sync.RWMutex, akv.Cache) {
	var h maphash.Hash
	i := h.Sum64() % n
	return c.rwmutexs[i], c.caches[i]
}

// Keys group

func (c *Cache) Del(ctx context.Context, keys []string) int {
	count := 0
	for _, key := range keys {
		rwmutex, cache := c.rwmutexcache(key)
		rwmutex.Lock()
		count += cache.Del(ctx, []string{key})
		rwmutex.Unlock()
	}
	return count
}

// Strings group

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.Get(ctx, key)
}

func (c *Cache) Set(ctx context.Context, key string, val []byte) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	cache.Set(ctx, key, val)
}
