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
	var ctx = req.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctxNew, cancel := context.WithTimeout(ctx, timeout)
	return req.Clone(ctxNew), cancel
}
