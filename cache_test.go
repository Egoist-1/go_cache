package go_cache

import (
	"context"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCache(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	err := client.Set(context.Background(), "key1", "aa", 0).Err()
	require.NoError(t, err)
	s := client.Decr(context.Background(), "key1").String()
	fmt.Println(s)
	lru.New()
}
