package go_cache

import (
	"errors"
	"fmt"
)

func assertFail(key string, typ string) error {
	return errors.New(fmt.Sprintf("assert fail %s not %s", key, typ))
}
