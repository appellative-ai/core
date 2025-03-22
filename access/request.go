package access

import (
	"net/http"
)

// Request - request attributes interface for non HTTP traffic
type Request interface {
	Url() string
	Header() http.Header
	Method() string
	Protocol() string
}

// RequestImpl - non HTTP request attributes
type RequestImpl struct {
	Url      string
	Header   http.Header
	Method   string
	Protocol string
}
