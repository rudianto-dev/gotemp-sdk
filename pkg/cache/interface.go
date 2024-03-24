package cache

import (
	"context"
	"time"
)

type IDatabase interface {
	Save(ctx context.Context, key string, payload interface{}, ttl time.Duration) (err error)
	Get(ctx context.Context, query string, params map[string]interface{}) (err error)
}
