package go_cache

import (
	"context"
	"sync"
)

type Cache interface {
	Set(ctx context.Context, key string, val any) error
	Get(ctx context.Context, key string) Result
	Delete(ctx context.Context, key string) Result
}

type Lcache struct {
	cache map[string]any
	mutex sync.RWMutex
	close sync.Once
}

func NewCache() Cache {
	cache := make(map[string]any, 256)
	return &Lcache{
		cache: cache,
		mutex: sync.RWMutex{},
		close: sync.Once{},
	}
}

func (m *Lcache) Set(ctx context.Context, key string, val any) error {
	//TODO implement me
	panic("implement me")
}

func (m *Lcache) Get(ctx context.Context, key string) Result {
	//TODO implement me
	panic("implement me")
}

func (m *Lcache) Delete(ctx context.Context, key string) Result {
	//TODO implement me
	panic("implement me")
}
func (m *Lcache) Close() error {
	m.close.Do(func() {
		m.cache = nil
	})
	return nil
}
