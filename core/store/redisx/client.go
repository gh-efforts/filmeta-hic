package redisx

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bitrainforest/filmeta-hic/core/assert"
	"github.com/go-redis/redis/v7"
)

type (
	Conf struct {
		Addr     string `json:"addr"`
		DB       int    `json:"db"`
		Password string `json:"password"`
		UserName string `json:"userName"`

		PoolSize    int `json:"pool_size"`
		MaxRetries  int `json:"max_retries"`
		IdleTimeout int `json:"idle_timeout"`
	}

	RedisNode struct {
		*redis.Client
	}
	RedisNodes []RedisNode
)

func MustInit(cfg *Conf) RedisNodes {
	nodes, err := Init(cfg)
	if err != nil {
		assert.CheckErr(fmt.Errorf("setup redis err: %s", err))
	}

	return nodes
}

func Init(cfg *Conf) (RedisNodes, error) {
	addrs := strings.Split(cfg.Addr, ",")
	if len(addrs) == 0 {
		return nil, fmt.Errorf("empty redis addr")
	}

	var nodes RedisNodes
	for _, addr := range addrs {
		if addr == "" {
			continue
		}
		cli := redis.NewClient(&redis.Options{
			Addr:        addr,
			DB:          cfg.DB,
			Password:    cfg.Password,
			Username:    cfg.UserName,
			MaxRetries:  cfg.MaxRetries,
			PoolSize:    cfg.PoolSize,
			IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		})

		//cli.AddHook()

		ping := cli.Ping()
		if ping.Val() == "" {
			return nil, fmt.Errorf("redis connect addr %s err: %s", addr, ping.Err())
		}

		nodes = append(nodes, RedisNode{cli})
	}

	return nodes, nil
}

func (nodes RedisNodes) GetClient(ctx context.Context) RedisNode {
	return RedisNode{
		RoundRobin(nodes).WithContext(ctx),
	}
}

func (node RedisNode) NewLocker(key string, opts ...RedisLockOption) *RedisLock {
	return NewRedisLock(node.Client, key, opts...)
}
