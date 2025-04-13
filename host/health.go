package host

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/rest"
	"net/http"
)

var (
	HealthEndpoint = rest.NewEndpoint(ExchangeHandler, nil, func(r *http.Request) (*http.Response, error) {
		return httpx.NewResponse(http.StatusOK, nil, []byte("up")), nil
	})
)
