package example

import (
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
	"unsafe"
)

func getSize[T any]() uintptr {
	var v T
	return unsafe.Sizeof(v)
}

func TestName(t *testing.T) {
	mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	mc.Decrement()
	it, err := mc.Get("foo")
	mc.
}
