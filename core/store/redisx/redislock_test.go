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
	firstLock.SetExpire(30)
	firstAcquire, err := firstLock.Acquire(ctx, key)
	assert.Nil(t, err)
	assert.True(t, firstAcquire)

	secondLock := NewRedisLock(client)

	secondLock.SetExpire(30)
	againAcquire, err := secondLock.Acquire(ctx, key)
	assert.Nil(t, err)
	assert.False(t, againAcquire)

	release := firstLock.Release(ctx, key)
	assert.True(t, release)

	endAcquire, err := secondLock.Acquire(ctx, key)
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedisLock(tt.args.store, tt.args.options...)
			assert.Equal(t, got.seconds, tt.want.seconds)
			assert.Nil(t, got.store)
		})
	}
}
