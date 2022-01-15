package internal

import (
	"context"

	"github.com/decimalbell/akv"
)

type Cache struct {
	m map[string]*Value
}

func NewCache() akv.Cache {
	return &Cache{
		m: make(map[string]*Value),
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

func (c *Cache) HGetAll(ctx context.Context, key string) ([][]byte, error) {
	h, err := c.getHashValue(ctx, key)
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, nil
	}
	return h.items(), nil
}

func (c *Cache) HKeys(ctx context.Context, key string) ([]string, error) {
	h, err := c.getHashValue(ctx, key)
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, nil
	}
	return h.keys(), nil
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

func (c *Cache) HVals(ctx context.Context, key string) ([][]byte, error) {
	h, err := c.getHashValue(ctx, key)
	if err != nil {
		return nil, err
	}
	if h == nil {
		return nil, nil
	}
	return h.values(), nil
}

// Keys group

func (c *Cache) getSetValue(ctx context.Context, key string) (Set, error) {
	value, ok := c.m[key]
	if !ok {
		return nil, nil
	}
	return value.set()
}

func (c *Cache) setDefaultSetValue(ctx context.Context, key string) (Set, error) {
	value, ok := c.m[key]
	if !ok {
		value = newSetValue()
		c.m[key] = value
	}
	return value.set()
}

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

// Sets group

func (c *Cache) SAdd(ctx context.Context, key string, members []string) (int, error) {
	s, err := c.setDefaultSetValue(ctx, key)
	if err != nil {
		return 0, err
	}
	return s.add(members), nil
}

func (c *Cache) SCard(ctx context.Context, key string) (int, error) {
	s, err := c.getSetValue(ctx, key)
	if err != nil {
		return 0, err
	}
	if s == nil {
		return 0, nil
	}
	return s.len(), nil
}

func (c *Cache) SIsMember(ctx context.Context, key string, member string) (int, error) {
	s, err := c.getSetValue(ctx, key)
	if err != nil {
		return 0, err
	}
	if s == nil {
		return 0, nil
	}

	if s.contains(member) {
		return 1, nil
	}
	return 0, nil
}

func (c *Cache) SMIsMember(ctx context.Context, key string, members []string) ([]int, error) {
	s, err := c.getSetValue(ctx, key)
	if err != nil {
		return nil, err
	}
	results := make([]int, len(members))
	if s == nil {
		return results, nil
	}

	for i, member := range members {
		result := 0
		if s.contains(member) {
			result = 1
		}
		results[i] = result

	}
	return results, nil
}

func (c *Cache) SRem(ctx context.Context, key string, members []string) (int, error) {
	s, err := c.getSetValue(ctx, key)
	if err != nil {
		return 0, err
	}
	if s == nil {
		return 0, nil
	}
	return s.remove(members), nil
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
