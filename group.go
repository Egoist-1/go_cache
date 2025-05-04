package go_cache

import "sync"

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64,
	getter Getter) *Group {
	if getter != nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes,
		},
	}
	groups[name] = g
	return g
}

func (g *Group) Get(key string) (ByteView, error) {

}
func (g *Group) load(key string) (value ByteView,
	err error) {
}

func (g *Group) getLocally(key string) (ByteView, error) {

}

func (g *Group) populateCache(key string, value ByteView) {

}
