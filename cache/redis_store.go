package cache

import "context"

type RedisStore struct {
}

func (r RedisStore) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}
