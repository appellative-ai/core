package httpx

import (
	"github.com/behavioral-ai/core/messaging"
	http2 "net/http"
)

type Agent interface {
	messaging.Agent
	Exchange(r *http2.Request, next *Frame) (*http2.Response, error)
}
