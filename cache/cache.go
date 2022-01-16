package cache

import (
	"context"
	"hash/fnv"
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
	h := fnv.New32a()
	i := h.Sum32() % n
	return c.rwmutexs[i], c.caches[i]
}

// Hashes group

func (c *Cache) HDel(ctx context.Context, key string, fields []string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	return cache.HDel(ctx, key, fields)
}

func (c *Cache) HGet(ctx context.Context, key string, field string) ([]byte, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.HGet(ctx, key, field)
}

func (c *Cache) HGetAll(ctx context.Context, key string) ([][]byte, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.HGetAll(ctx, key)
}

func (c *Cache) HKeys(ctx context.Context, key string) ([]string, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.HKeys(ctx, key)
}

func (c *Cache) HLen(ctx context.Context, key string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.HLen(ctx, key)
}

func (c *Cache) HSet(ctx context.Context, key string, field string, val []byte) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	return cache.HSet(ctx, key, field, val)
}

func (c *Cache) HStrLen(ctx context.Context, key string, field string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	return cache.HStrLen(ctx, key, field)
}

func (c *Cache) HVals(ctx context.Context, key string) ([][]byte, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.HVals(ctx, key)
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

// Sets group

func (c *Cache) SAdd(ctx context.Context, key string, members []string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	return cache.SAdd(ctx, key, members)
}

func (c *Cache) SCard(ctx context.Context, key string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.SCard(ctx, key)
}

func (c *Cache) SIsMember(ctx context.Context, key string, member string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.SIsMember(ctx, key, member)
}

func (c *Cache) SMIsMember(ctx context.Context, key string, members []string) ([]int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return cache.SMIsMember(ctx, key, members)
}

func (c *Cache) SRem(ctx context.Context, key string, members []string) (int, error) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	return cache.SRem(ctx, key, members)
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
