package httpx

import (
	"context"
	"time"
)

var (
	cancelFn = func() {}
)

func NewContext(timeout time.Duration) (context.Context, func()) {
	if timeout > 0 {
		return context.WithTimeout(context.Background(), timeout)
	}
	return context.Background(), cancelFn
}
