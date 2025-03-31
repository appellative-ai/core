package httpx

import (
	"github.com/behavioral-ai/core/messaging"
)

func NewConfigExchangeMessage(ex Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, messaging.ConfigEvent)
	m.SetContent(messaging.ContentTypeExchange, ex)
	return m
}

func ConfigExchangeContent(m *messaging.Message) Exchange {
	if m.Event() != messaging.ConfigEvent || m.ContentType() != messaging.ContentTypeExchange {
		return nil
	}
	if cfg, ok := m.Body.(Exchange); ok {
		return cfg
	}
	return nil
}
