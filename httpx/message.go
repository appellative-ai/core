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

func ConfigExchangeContent(m *messaging.Message) Exchange {
	if m.Event() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchange {
		return nil
	}
	if cfg, ok := m.Body.(Exchange); ok {
		return cfg
	}
	return nil
}

func NewConfigExchangeWriterMessage(ex ExchangeWriter) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchangeWriter, ex)
	return m
}

func ConfigExchangeWriterContent(m *messaging.Message) ExchangeWriter {
	if m.Event() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchangeWriter {
		return nil
	}
	if cfg, ok := m.Body.(ExchangeWriter); ok {
		return cfg
	}
	return nil
}
