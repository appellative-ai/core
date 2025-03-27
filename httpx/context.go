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
	ctxNew, cancel := NewContext(req, timeout)
	return req.Clone(ctxNew), cancel
}

func NewContext(req *http.Request, timeout time.Duration) (context.Context, func()) {
	var ctx = context.Background()
	if req != nil && req.Context() != nil {
		ctx = req.Context()
	}
	if timeout <= 0 || req == nil {
		return ctx, cancelFn
	}
	return context.WithTimeout(ctx, timeout)
}
