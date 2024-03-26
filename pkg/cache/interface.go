package cache

import (
	"context"
	"time"
)

type Service interface {
	Get(ctx context.Context, key string) (res string, err error)
	Save(ctx context.Context, key string, payload interface{}, ttl time.Duration) (err error)
	Delete(ctx context.Context, key string) (err error)
}
