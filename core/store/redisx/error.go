package redisx

import (
	"context"
	"fmt"
)

var (
	RedisErr = fmt.Errorf("internal cache errors")
)

func WrapErr(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	//todo log
	return RedisErr
}
