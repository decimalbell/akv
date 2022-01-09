package akv

import (
	"context"
	"errors"
)

var (
	ErrWrongType = errors.New("akv: wrong type")
)

type Cache interface {
	// Keys group
	Del(ctx context.Context, keys []string) int

	// Strings group
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, val []byte)
}
