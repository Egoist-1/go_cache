package example

import (
	"context"
	"fmt"
	"go_cache"
	"testing"
)

func TestExample_cache(t *testing.T) {
	cache := go_cache.NewCache(go_cache.Option{
		MaxCap:    0,
		OnEvicted: nil,
	})
	err := cache.Set(context.Background(), "key1", 1, 0)
	if err != nil {

	}
	i, err := cache.Get(context.Background(), "key1").Int()
	if err != nil {

	}
	fmt.Println(i)
}
