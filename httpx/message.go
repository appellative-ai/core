package httpx

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
)

const (
	ContentTypeExchange = "application/exchange"
	DefaultRelatesTo    = "default"
)

func NewConfigExchangeMessage(ex rest.Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchange, ex)
	return m
}

func ConfigExchangeContent(m *messaging.Message) (rest.Exchange, bool) {
	if m.Name() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchange {
		return nil, false
	}
	if cfg, ok := m.Body.(rest.Exchange); ok {
		return cfg, true
	}
	return nil, false
}

/*

	ContentTypeExchangeWriter = "application/exchange-writer"

func NewConfigExchangeWriterMessage(ex ExchangeWriter) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchangeWriter, ex)
	return m
}

func ConfigExchangeWriterContent(m *messaging.Message) (ExchangeWriter, bool) {
	if m.Name() != messaging.ConfigEvent || m.ContentType() != ContentTypeExchangeWriter {
		return nil, false
	}
	if cfg, ok := m.Body.(ExchangeWriter); ok {
		return cfg, true
	}
	return nil, false
}


*/
