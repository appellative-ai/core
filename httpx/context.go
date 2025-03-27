package httpx

import (
	"context"
	"net/http"
	"time"
)

var (
	cancelFn = func() {}
)

func NewRequestWithTimeout(req *http.Request, timeout time.Duration) (*http.Request, func()) {
	if timeout <= 0 || req == nil {
		return req, cancelFn
	}
	ctxNew, cancel := NewContext(req.Context(), timeout)
	return req.Clone(ctxNew), cancel
}

func NewContext(ctx context.Context, timeout time.Duration) (context.Context, func()) {
	if ctx == nil {
		ctx = context.Background()
	}
	if timeout <= 0 {
		return ctx, cancelFn
	}
	return context.WithTimeout(ctx, timeout)
}
