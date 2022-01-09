package internal

import (
	"context"

	"github.com/decimalbell/akv"
)

type Cache struct {
	m map[string]*value
}

func NewCache() akv.Cache {
	return &Cache{
		m: make(map[string]*value),
	}
}

// Keys group

func (c *Cache) Del(ctx context.Context, keys []string) int {
	count := 0
	for _, key := range keys {
		_, ok := c.m[key]
		if ok {
			delete(c.m, key)
			count++
		}
	}
	return count
}

// Strings group

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	value, ok := c.m[key]
	if !ok {
		return nil, nil
	}
	return value.bytes()
}

func (c *Cache) Set(ctx context.Context, key string, val []byte) {
	c.m[key] = newBytesValue(val)
}
