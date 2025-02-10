package host

import (
	"github.com/behavioral-ai/core/messaging"
)

func Ping(uri any) *aspect.Status {
	return messaging.Ping(Exchange, uri)
}
