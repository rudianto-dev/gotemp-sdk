package cache

import (
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"

	"github.com/go-redis/redis"
)

type DataSource struct {
	Redis  *redis.Client
	Logger *logger.Logger
}

func NewCache(redisConfig *redis.Options, logger *logger.Logger) (ds *DataSource, err error) {
	ds = &DataSource{Logger: logger}
	err = ds.ConnectRedisClient(redisConfig)
	if err != nil {
		return
	}
	return
}

func (ds *DataSource) ConnectRedisClient(cf *redis.Options) (err error) {
	ds.Redis = redis.NewClient(cf)
	if _, err = ds.Redis.Ping().Result(); err != nil {
		return
	}
	return
}

func (ds *DataSource) CloseConnection() (err error) {
	if ds.Redis != nil {
		err = ds.Redis.Close()
		return
	}
	return
}
