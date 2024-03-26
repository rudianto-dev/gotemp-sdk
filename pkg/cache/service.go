package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

func (ds *DataSource) Get(ctx context.Context, key string) (res string, err error) {
	res, err = ds.Redis.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return
	}
	return
}

func (ds *DataSource) Save(ctx context.Context, key string, payload interface{}, ttl time.Duration) (err error) {
	_, err = ds.Redis.Set(key, payload, ttl).Result()
	return
}

func (ds *DataSource) Delete(ctx context.Context, key string) (err error) {
	_, err = ds.Redis.Del(key).Result()
	return
}
