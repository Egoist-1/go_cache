package go_cache

// ByteView 只读数据结构,用来表示缓存之
type ByteView struct {
	b []byte
}

func (b ByteView) Len() int {
	return len(b.b)
}

func (b ByteView) ByteSlice() []byte {
	return cloneBytes(b.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
