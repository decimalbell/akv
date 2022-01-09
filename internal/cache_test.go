package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	value := []byte("value")

	cache.Set(ctx, key, value)
	v, err := cache.Get(ctx, key)
	assert.Equal(t, err, nil)
	assert.Equal(t, value, v)
}

func TestDel(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	value := []byte("value")

	assert.Equal(t, 0, cache.Del(ctx, []string{key}))

	cache.Set(ctx, key, value)
	assert.Equal(t, 1, cache.Del(ctx, []string{key}))
}
