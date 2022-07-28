package store

import (
	"context"

	"github.com/bitrainforest/filmeta-hic/core/assert"

	configHelper "github.com/bitrainforest/filmeta-hic/core/config"

	"github.com/bitrainforest/filmeta-hic/core/store/redisx"
	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/config"
)

var (
	redisBundle redisx.RedisNodes
)

func GetRedisClient(ctx context.Context) *redis.Client {
	return redisBundle.GetClient(ctx).Client
}

func MustLoadRedis(confSource config.Config, key string) {
	var (
		conf redisx.Conf
	)
	if err := configHelper.ScanConfValue(confSource, key, &conf); err != nil {
		assert.CheckErr(err)
	}
	redisBundle = redisx.MustInit(&conf)
}
