package host

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
)

func Ping(uri any) *aspect.Status {
	return messagingx.Ping(Exchange, uri)
}
