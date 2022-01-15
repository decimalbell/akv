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

// Hashes group

func (c *Cache) getHashValue(ctx context.Context, key string) (Map, error) {
	value, ok := c.m[key]
	if !ok {
		return nil, nil
	}
	return value.hash()
}

func (c *Cache) setDefaultHashValue(ctx context.Context, key string) (Map, error) {
	value, ok := c.m[key]
	if !ok {
		value = newHashValue()
		c.m[key] = value
	}
	return value.hash()
}

func (c *Cache) HDel(ctx context.Context, key string, fields []string) (int, error) {
	h, err := c.getHashValue(ctx, key)
	if err != nil {
		return 0, err
	}
	if h == nil {
		return 0, nil
	}

	count := 0
	for _, field := range fields {
		_, ok := h[field]
		if ok {
			delete(h, field)
			count++
		}
	}

	return count, nil
}

func (c *Cache) HGet(ctx context.Context, key string, field string) ([]byte, error) {
	h, err := c.getHashValue(ctx, key)
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, nil
	}
	return h[field], nil
}

func (c *Cache) HSet(ctx context.Context, key string, field string, val []byte) (int, error) {
	h, err := c.setDefaultHashValue(ctx, key)
	if err != nil {
		return 0, err
	}

	if _, ok := h[field]; ok {
		h[field] = val
		return 0, nil
	}
	h[field] = val
	return 1, nil
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
