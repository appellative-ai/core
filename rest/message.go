package rest

import "github.com/behavioral-ai/core/messaging"

const (
	ContentTypeRoute    = "application/x-route"
	ContentTypeExchange = "application/x-exchange"
)

func NewRouteMessage(name, uri string, ex Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	r := NewRoute(name, uri, ex)
	m.SetContent(ContentTypeRoute, "", r)
	return m
}

func RouteContent(m *messaging.Message) (*Route, bool) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeRoute) {
		return nil, false
	}
	if v, ok := m.Content.Value.(*Route); ok {
		return v, true
	}
	return nil, false
}

func NewExchangeMessage(ex Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeExchange, "", ex)
	return m
}

func ExchangeContent(m *messaging.Message) (Exchange, bool) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeExchange) {
		return nil, false
	}
	if cfg, ok := m.Content.Value.(Exchange); ok {
		return cfg, true
	}
	return nil, false
}
