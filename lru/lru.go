package lru

import "container/list"

func New(maxBytes int64, onEvicted func(k string, v Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[interface{}]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Cache 并发访问并不安全
type Cache struct {
	maxBytes int64 //允许最大的内存
	nowBytes int64 //目前使用的内存
	ll       *list.List
	//key 是字符串 val 是双向链表中对应节点的指针
	cache     map[interface{}]*list.Element
	OnEvicted func(k string, v Value) //删除后的回调函数
}

func (c *Cache) Get(key string) (val Value, ok bool) {
	ele, ok := c.cache[key]
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除 移除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.ll.Remove(ele)
		c.nowBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add
func (c *Cache) Add(key string, val Value) {
	//key存在 update
	if ele, ok := c.cache[key]; ok {
		kv := ele.Value.(*entry)
		c.ll.MoveToFront(ele)
		c.nowBytes += int64(val.Len()) - int64(kv.value.Len())
		kv.value = val
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: val})
		c.cache[key] = ele
		c.nowBytes += int64(len(key)) + int64(val.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nowBytes {
		c.RemoveOldest()
	}
}

// Len 获取链表中有多少条数据
func (c *Cache) Len() int {
	return c.ll.Len()
}

// entry 双向链表的数据类型
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}
