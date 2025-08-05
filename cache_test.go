package go_cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache(t *testing.T) {
	str := "hello world"
	fmt.Println(len(str))
}

func TestCacheSize(t *testing.T) {
	testCase := []struct {
		name   string
		input  any
		output uintptr
	}{
		{
			name:   "int",
			input:  int64(10),
			output: 8,
		},
		{},
		{
			name:   "string", //struct{addr uintptr,len int} //存储字符串的结构体占用16个字节
			input:  "hello world",
			output: 16 + 11,
		},
		{
			name: "struct",
			input: struct {
				age int
				str string
			}{
				age: 231,
				str: "22",
			},
			output: 16 + 2 + 8,
		},
		{
			name: "嵌套结构体",
			input: struct {
				name string
				age  int
				s    struct {
					name2 string
				}
			}{
				name: "hello",
				age:  2,
				s: struct{ name2 string }{
					name2: "world",
				},
			},
			output: 16*2 + 5 + 5 + 8,
		},
	}
	cache := NewCache(Option{
		MaxCap:    100,
		OnEvicted: nil,
	})
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			size := cache.size(tc.input)
			fmt.Println(size)
			assert.Equal(t, tc.output, size)
		})
	}
}
