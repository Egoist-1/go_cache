package go_cache

import "time"

type Result struct {
	Err error
	v   val
}
type val struct {
	value      any
	expiration time.Duration
}
