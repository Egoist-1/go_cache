package go_cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var (
	AssertFail = errors.New("assert fail")
)

type item struct {
	key        string
	val        any
	expiration time.Duration
}

type Result struct {
	Err error
	ctx context.Context
	item
}

// Bind 请确保set时是已经序列化好的
func (r *Result) Bind(a any) (err error) {
	if r.Err != nil {
		return r.Err
	}
	bytes, ok := r.val.([]byte)
	if !ok {
		return ErrTypeFail
	}
	return json.Unmarshal(bytes, a)
}

func (r *Result) Sting() (str string, err error) {
	if r.Err != nil {
		return "", r.Err
	}
	str, ok := r.val.(string)
	if !ok {
		return "", assertFail(r.key, "string")
	}
	return str, nil
}

func (r *Result) Any() (any, error) {
	return r.item, r.Err
}
func (r *Result) Int() (int, error) {
	if r.Err != nil {
		return 0, r.Err
	}
	i, ok := r.val.(int)
	if !ok {
		return 0, ErrTypeFail
	}
	return i, nil
}
