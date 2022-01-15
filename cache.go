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
	HGetAll(ctx context.Context, key string) ([][]byte, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HLen(ctx context.Context, key string) (int, error)
	HSet(ctx context.Context, key string, field string, val []byte) (int, error)
	HVals(ctx context.Context, key string) ([][]byte, error)

	// Keys group
	Del(ctx context.Context, keys []string) int

	// Sets group
	SAdd(ctx context.Context, key string, members []string) (int, error)
	SCard(ctx context.Context, key string) (int, error)
	SIsMember(ctx context.Context, key string, member string) (int, error)
	SMIsMember(ctx context.Context, key string, members []string) ([]int, error)
	SRem(ctx context.Context, key string, members []string) (int, error)

	// Strings group
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, val []byte)
}
