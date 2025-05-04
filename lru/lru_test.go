package lru

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type str string

func (s str) Len() int {
	return 10
}

func TestLru(t *testing.T) {
	testCase := []struct {
		name    string
		cache   *Cache
		wantLen int64
		before  func(t *testing.T, cache *Cache)
		after   func(t *testing.T, cache *Cache)
	}{
		{
			name:    "当内存超过设定值时,是否会触发节点移除",
			cache:   New(100, nil),
			wantLen: 95,
			before: func(t *testing.T, cache *Cache) {
				cache.nowBytes = 85
				cache.Add("k1", str("1"))
				cache.Add("k2", str("2"))
			},
			after: func(t *testing.T, cache *Cache) {
				val, ok := cache.Get("k2")
				assert.Equal(t, ok, true)
				assert.Equal(t, val, str("2"))
				_, ok = cache.Get("k1")
				assert.Equal(t, ok, false)

			},
		},
		{
			name: "回调函数是否被调用",
			cache: New(100, func(k string, v Value) {
				fmt.Println(k, v)
			}),
			wantLen: 95,
			before: func(t *testing.T, cache *Cache) {
				cache.nowBytes = 85
				cache.Add("k1", str("1"))
				cache.Add("k2", str("2"))
			},
			after: func(t *testing.T, cache *Cache) {
				val, ok := cache.Get("k2")
				assert.Equal(t, ok, true)
				assert.Equal(t, val, str("2"))
				_, ok = cache.Get("k1")
				assert.Equal(t, ok, false)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t, tc.cache)
			tc.after(t, tc.cache)
		})
	}
}
