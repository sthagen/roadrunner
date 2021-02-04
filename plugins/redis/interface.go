package redis

import "github.com/go-redis/redis/v8"

// Redis in the redis KV plugin interface
type Redis interface {
	// GetClient
	GetClient() redis.UniversalClient
}
