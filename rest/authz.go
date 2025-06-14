package rest

import (
	"net/http"
)

const (
	NamespaceNameAuth = "test:resiliency:link/authorization/http"
	AuthorizationName = "Authorization"
)

func Authorization(next Exchange) Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(AuthorizationName)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		return next(r)
	}
}
