package akv

import (
	"context"
	"errors"
)

var (
	ErrWrongType = errors.New("akv: wrong type")
)

type Cache interface {
	// Hashes group
	HDel(ctx context.Context, key string, fields []string) (int, error)
	HGet(ctx context.Context, key string, field string) ([]byte, error)
	HSet(ctx context.Context, key string, field string, val []byte) (int, error)

	// Keys group
	Del(ctx context.Context, keys []string) int

	// Sets group
	SAdd(ctx context.Context, key string, members []string) (int, error)
	SCard(ctx context.Context, key string) (int, error)
	SRem(ctx context.Context, key string, members []string) (int, error)

	// Strings group
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, val []byte)
}
