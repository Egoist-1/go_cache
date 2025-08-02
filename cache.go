package go_cache

import (
	"container/list"
	"context"
	"errors"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

const (
	KB int64 = 1024
	MB int64 = 1024 * KB
	GB int64 = 1024 * MB
)

var (
	ErrTypeFail = errors.New("type fail")
	NotFond     = errors.New("cache: not fond")
)

type Cache interface {
	Set(ctx context.Context, key string, val any, expiration time.Duration) error
	Get(ctx context.Context, key string) *Result
	Delete(ctx context.Context, key string) error
}

func NewCache(opt Option) *Lcache {
	cache := &Lcache{
		list:  list.New(),
		cache: make(map[string]*list.Element),
		lock:  sync.RWMutex{},
		close: sync.Once{},
		nCap:  0,
		opt:   opt,
	}
	return cache
}

type Lcache struct {
	list  *list.List
	cache map[string]*list.Element
	lock  sync.RWMutex
	close sync.Once
	nCap  int
	opt   Option
}

type Option struct {
	MaxCap    int
	OnEvicted func(key string, value any)
}

func (m *Lcache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	v := item{
		val:        val,
		expiration: expiration,
	}
	m.list.PushFront(v)
	//TODO implement me
	panic("implement me")
}

func (m *Lcache) Get(ctx context.Context, key string) (res *Result) {
	res.ctx = ctx
	m.lock.RLock()
	var ele *list.Element
	ele, ok := m.cache[key]
	if !ok {
		res.Err = NotFond
		return
	}
	m.lock.RUnlock()
	val := ele.Value.(item)
	if val.expiration > 0 {
		m.opt.OnEvicted(val.key, val.val)
		return
	}

	m.list.MoveToFront(ele)
	res.item = val
	return

}

func (m *Lcache) Delete(ctx context.Context, key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	val := m.list.Remove(m.cache[key]).(item)
	delete(m.cache, key)
	m.opt.OnEvicted(val.key, val.val)
	return nil
}

func (m *Lcache) Incr(ctx context.Context, key string) (res *Result) {
	m.lock.Lock()
	defer m.lock.Unlock()
	return
}

func (m *Lcache) Decr(ctx context.Context, key string) Result {
	//TODO implement me
	panic("implement me")
}

func (m *Lcache) size(a any) uintptr {
	return sizeOfRecursive(a)
}

func sizeOfRecursive(v any) uintptr {
	visited := make(map[uintptr]bool)
	return sizeOf(reflect.ValueOf(v), visited)
}

func sizeOf(val reflect.Value, visited map[uintptr]bool) uintptr {
	if !val.IsValid() {
		return 0
	}

	typ := val.Type()
	kind := typ.Kind()

	// If pointer, dereference and avoid double-counting
	if kind == reflect.Ptr || kind == reflect.Interface {
		ptr := val.Pointer()
		if ptr == 0 || visited[ptr] {
			return 0
		}
		visited[ptr] = true
		return unsafe.Sizeof(ptr) + sizeOf(val.Elem(), visited)
	}

	switch kind {
	case reflect.Array:
		var total uintptr
		for i := 0; i < val.Len(); i++ {
			total += sizeOf(val.Index(i), visited)
		}
		return total
	case reflect.Slice:
		if val.IsNil() {
			return unsafe.Sizeof(val.Interface()) // header only
		}
		var total uintptr = unsafe.Sizeof(val.Interface()) // slice header
		for i := 0; i < val.Len(); i++ {
			total += sizeOf(val.Index(i), visited)
		}
		return total
	case reflect.Map:
		if val.IsNil() {
			return unsafe.Sizeof(val.Interface()) // map header only
		}
		var total uintptr = unsafe.Sizeof(val.Interface())
		for _, key := range val.MapKeys() {
			total += sizeOf(key, visited)
			total += sizeOf(val.MapIndex(key), visited)
		}
		return total
	case reflect.Struct:
		var total uintptr
		for i := 0; i < val.NumField(); i++ {
			total += sizeOf(val.Field(i), visited)
		}
		return total
	case reflect.String:
		return unsafe.Sizeof("") + uintptr(val.Len())
	default:
		return val.Type().Size()
	}
}
func (m *Lcache) Close() error {
	m.close.Do(func() {
		m.cache = nil
	})
	return nil
}
