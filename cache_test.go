package go_cache

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestCache(t *testing.T) {
	client := redis.NewClient(nil)
	client.Set()
	get := client.Get(nil, "")
}
