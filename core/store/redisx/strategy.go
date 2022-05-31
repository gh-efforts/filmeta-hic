package redisx

import (
	"sync/atomic"
)

var (
	roundRobinCount uint64
)

// RoundRobin  is a round-robin load balancing strategy.
func RoundRobin(clients []RedisNode) RedisNode {
	if len(clients) == 1 {
		return clients[0]
	}
	i := (roundRobinCount) % uint64(len(clients))
	atomic.AddUint64(&roundRobinCount, 1)

	return clients[i]
}
