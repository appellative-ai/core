package httpx

import (
	"github.com/behavioral-ai/core/messaging"
)

// Agent - adds a chainable exchange method
type Agent interface {
	messaging.Agent
	Exchange(next Exchange) Exchange
}
