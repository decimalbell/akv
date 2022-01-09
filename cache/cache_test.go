package cache

import (
	"context"
	"strconv"
	"testing"
)

func BenchmarkDel(b *testing.B) {
	cache := New()
	ctx := context.TODO()

	keys := make([]string, b.N)
	values := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = strconv.Itoa(i)
		values[i] = []byte(strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		cache.Set(ctx, keys[i], values[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del(ctx, []string{keys[i]})
	}
}

func BenchmarkGet(b *testing.B) {
	cache := New()
	ctx := context.TODO()

	keys := make([]string, b.N)
	values := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = strconv.Itoa(i)
		values[i] = []byte(strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		cache.Set(ctx, keys[i], values[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(ctx, keys[i])
	}
}

func BenchmarkSet(b *testing.B) {
	cache := New()
	ctx := context.TODO()

	keys := make([]string, b.N)
	values := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = strconv.Itoa(i)
		values[i] = []byte(strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(ctx, keys[i], values[i])
	}
}
