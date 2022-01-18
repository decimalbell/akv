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

func (c *Cache) withLock(key string, fn func(cache akv.Cache)) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.Lock()
	defer rwmutex.Unlock()

	fn(cache)
}

func (c *Cache) withRLock(key string, fn func(cache akv.Cache)) {
	rwmutex, cache := c.rwmutexcache(key)
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	fn(cache)
}

// Hashes group

func (c *Cache) HDel(ctx context.Context, key string, fields []string) (n int, err error) {
	c.withLock(key, func(cache akv.Cache) {
		n, err = cache.HDel(ctx, key, fields)
	})
	return n, err
}

func (c *Cache) HGet(ctx context.Context, key string, field string) (value []byte, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		value, err = cache.HGet(ctx, key, field)
	})
	return value, err
}

func (c *Cache) HGetAll(ctx context.Context, key string) (kvs [][]byte, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		kvs, err = cache.HGetAll(ctx, key)
	})
	return kvs, err
}

func (c *Cache) HKeys(ctx context.Context, key string) (keys []string, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		keys, err = cache.HKeys(ctx, key)
	})
	return keys, err
}

func (c *Cache) HLen(ctx context.Context, key string) (n int, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		n, err = cache.HLen(ctx, key)
	})
	return n, err
}

func (c *Cache) HSet(ctx context.Context, key string, field string, val []byte) (n int, err error) {
	c.withLock(key, func(cache akv.Cache) {
		n, err = cache.HSet(ctx, key, field, val)
	})
	return n, err
}

func (c *Cache) HStrLen(ctx context.Context, key string, field string) (n int, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		n, err = cache.HStrLen(ctx, key, field)
	})
	return n, err
}

func (c *Cache) HVals(ctx context.Context, key string) (values [][]byte, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		values, err = cache.HVals(ctx, key)
	})
	return values, err
}

// Keys group

func (c *Cache) Del(ctx context.Context, keys []string) int {
	count := 0
	for _, key := range keys {
		c.withRLock(key, func(cache akv.Cache) {
			count += cache.Del(ctx, []string{key})
		})
	}
	return count
}

// Sets group

func (c *Cache) SAdd(ctx context.Context, key string, members []string) (n int, err error) {
	c.withLock(key, func(cache akv.Cache) {
		n, err = cache.SAdd(ctx, key, members)
	})
	return n, err
}

func (c *Cache) SCard(ctx context.Context, key string) (n int, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		n, err = cache.SCard(ctx, key)
	})
	return n, err
}

func (c *Cache) SIsMember(ctx context.Context, key string, member string) (n int, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		n, err = cache.SIsMember(ctx, key, member)
	})
	return n, err
}

func (c *Cache) SMembers(ctx context.Context, key string) (members []string, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		members, err = cache.SMembers(ctx, key)
	})
	return members, err
}

func (c *Cache) SMIsMember(ctx context.Context, key string, members []string) (nums []int, err error) {
	c.withRLock(key, func(cache akv.Cache) {
		nums, err = cache.SMIsMember(ctx, key, members)
	})
	return nums, err
}

func (c *Cache) SRem(ctx context.Context, key string, members []string) (n int, err error) {
	c.withLock(key, func(cache akv.Cache) {
		n, err = cache.SRem(ctx, key, members)
	})
	return n, err
}

// Strings group

func (c *Cache) Get(ctx context.Context, key string) (value []byte, err error) {
	c.withLock(key, func(cache akv.Cache) {
		value, err = cache.Get(ctx, key)
	})
	return value, err
}

func (c *Cache) Set(ctx context.Context, key string, val []byte) {
	c.withLock(key, func(cache akv.Cache) {
		cache.Set(ctx, key, val)
	})
}
