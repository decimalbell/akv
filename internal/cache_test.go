package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hashes group
func TestHDel(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	{
		n, err := cache.HDel(ctx, key, []string{field})
		assert.Equal(t, 0, n)
		assert.Nil(t, err)
	}

	{
		cache.HSet(ctx, key, field, value)
		n, err := cache.HDel(ctx, key, []string{field})
		assert.Equal(t, 1, n)
		assert.Nil(t, err)
	}
}

func TestHGet(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	{
		v, err := cache.HGet(ctx, key, field)
		assert.Nil(t, err)
		assert.Nil(t, v)
	}

	{
		cache.HSet(ctx, key, field, value)
		v, err := cache.HGet(ctx, key, field)
		assert.Nil(t, err)
		assert.Equal(t, v, value)
	}
}

func TestHGetAll(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	cache.HSet(ctx, key, field, value)
	v, err := cache.HGetAll(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, [][]byte{[]byte(field), value}, v)
}

func TestHKeys(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	cache.HSet(ctx, key, field, value)
	v, err := cache.HKeys(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, []string{field}, v)
}

func TestHLen(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	{
		v, err := cache.HLen(ctx, key)
		assert.Nil(t, err)
		assert.Equal(t, 0, v)
	}

	{
		cache.HSet(ctx, key, field, value)
		v, err := cache.HLen(ctx, key)
		assert.Nil(t, err)
		assert.Equal(t, 1, v)
	}
}

func TestHSet(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	cache.HSet(ctx, key, field, value)
	v, err := cache.HGet(ctx, key, field)
	assert.Nil(t, err)
	assert.Equal(t, v, value)
}

func TestHVals(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	field := "field"
	value := []byte("value")

	cache.HSet(ctx, key, field, value)
	v, err := cache.HVals(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, [][]byte{value}, v)
}

// Keys group

func TestDel(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	value := []byte("value")

	assert.Equal(t, 0, cache.Del(ctx, []string{key}))

	cache.Set(ctx, key, value)
	assert.Equal(t, 1, cache.Del(ctx, []string{key}))
}

// Sets group

func TestSAdd(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	members := []string{"member"}

	{
		n, err := cache.SCard(ctx, key)
		assert.Nil(t, err)
		assert.Equal(t, 0, n)
	}

	{
		n, err := cache.SAdd(ctx, key, members)
		assert.Nil(t, err)
		assert.Equal(t, len(members), n)

		n, err = cache.SCard(ctx, key)
		assert.Nil(t, err)
		assert.Equal(t, len(members), n)
	}
}

func TestSCard(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	members := []string{"member"}

	n, err := cache.SCard(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)

	n, err = cache.SAdd(ctx, key, members)
	assert.Nil(t, err)
	assert.Equal(t, len(members), n)
	n, err = cache.SCard(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, len(members), n)

	n, err = cache.SRem(ctx, key, members)
	assert.Nil(t, err)
	assert.Equal(t, len(members), n)
	n, err = cache.SCard(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}

func TestSIsMember(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	member := "member"

	n, err := cache.SIsMember(ctx, key, member)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)

	n, err = cache.SAdd(ctx, key, []string{member})
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	n, err = cache.SIsMember(ctx, key, member)
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
}

func TestSMIsMember(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	members := []string{"member"}

	results, err := cache.SMIsMember(ctx, key, members)
	assert.Nil(t, err)
	assert.Equal(t, []int{0}, results)

	n, err := cache.SAdd(ctx, key, members)
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	results, err = cache.SMIsMember(ctx, key, members)
	assert.Nil(t, err)
	assert.Equal(t, []int{1}, results)
}

func TestSRem(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	members := []string{"member"}

	{
		n, err := cache.SRem(ctx, key, members)
		assert.Nil(t, err)
		assert.Equal(t, 0, n)
	}

	{
		n, err := cache.SAdd(ctx, key, members)
		assert.Nil(t, err)
		assert.Equal(t, len(members), n)

		n, err = cache.SRem(ctx, key, members)
		assert.Nil(t, err)
		assert.Equal(t, len(members), n)
	}
}

// Strings group

func TestGet(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	value := []byte("value")

	{
		v, err := cache.Get(ctx, key)
		assert.Nil(t, err)
		assert.Nil(t, v)
	}

	{
		cache.Set(ctx, key, value)
		v, err := cache.Get(ctx, key)
		assert.Nil(t, err)
		assert.Equal(t, value, v)
	}
}

func TestSet(t *testing.T) {
	cache := NewCache()
	ctx := context.TODO()
	key := "key"
	value := []byte("value")

	cache.Set(ctx, key, value)
	v, err := cache.Get(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, value, v)
}
