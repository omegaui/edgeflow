package cache

import "context"

type Interface interface {
	Get(ctx context.Context, key string) (string, error)
}
