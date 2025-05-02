package host

import "github.com/behavioral-ai/core/messaging"

var (
	exchange = messaging.NewExchange()
)

func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

func Agent(uri string) messaging.Agent {
	return exchange.Get(uri)
}

func Message(m *messaging.Message) {
	exchange.Message(m)
}

func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}
