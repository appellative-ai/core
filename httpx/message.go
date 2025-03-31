package httpx

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeExchange       = "application/exchange"
	ContentTypeExchangeWriter = "application/exchange-writer"
)

func NewConfigExchangeMessage(ex Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchange, ex)
	return m
}

func ConfigExchangeContent(m *messaging.Message) (Exchange, bool) {
	if m.Event() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchange {
		return nil, false
	}
	if cfg, ok := m.Body.(Exchange); ok {
		return cfg, true
	}
	return nil, false
}

func NewConfigExchangeWriterMessage(ex ExchangeWriter) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchangeWriter, ex)
	return m
}

func ConfigExchangeWriterContent(m *messaging.Message) (ExchangeWriter, bool) {
	if m.Event() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchangeWriter {
		return nil, false
	}
	if cfg, ok := m.Body.(ExchangeWriter); ok {
		return cfg, true
	}
	return nil, false
}
