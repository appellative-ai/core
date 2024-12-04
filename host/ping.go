package host

import (
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
)

func Ping(uri any) *core.Status {
	return messaging.Ping(Exchange, uri)
}
