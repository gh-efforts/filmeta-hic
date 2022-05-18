package redisx

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisLock(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	ctx := context.Background()
	key := "test"
	firstLock := NewRedisLock(client)
	firstLock.SetKey(key)
	firstLock.SetExpire(30)
	firstAcquire, err := firstLock.Acquire(ctx)
	assert.Nil(t, err)
	assert.True(t, firstAcquire)

	secondLock := NewRedisLock(client)
	secondLock.SetKey(key)

	secondLock.SetExpire(30)
	againAcquire, err := secondLock.Acquire(ctx)
	assert.Nil(t, err)
	assert.False(t, againAcquire)

	release := firstLock.Release(ctx)
	assert.True(t, release)

	endAcquire, err := secondLock.Acquire(ctx)
	assert.Nil(t, err)
	assert.True(t, endAcquire)
}

func TestNewRedisLock(t *testing.T) {
	type args struct {
		store   *redis.Client
		key     string
		options []RedisLockOption
	}
	tests := []struct {
		name string
		args args
		want *RedisLock
	}{
		{
			name: "test-second",
			args: args{
				store:   nil,
				key:     "default",
				options: []RedisLockOption{SetLockExpire(3)},
			},
			want: &RedisLock{
				store:   nil,
				seconds: 3,
				key:     "default",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedisLock(tt.args.store, tt.args.options...)
			assert.Equal(t, got.seconds, tt.want.seconds)
			assert.Equal(t, got.key, tt.want.key)
			assert.Nil(t, got.store)
		})
	}
}
