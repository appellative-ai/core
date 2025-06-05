package rest

import "github.com/behavioral-ai/core/messaging"

const (
	ContentTypeRoute    = "application/x-route"
	ContentTypeExchange = "application/x-exchange"
)

func NewRouteMessage(name, uri string, ex Exchange) *messaging.Message {
	return messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent).SetContent(ContentTypeRoute, "", NewRoute(name, uri, ex))
}

func RouteContent(m *messaging.Message) (*Route, *messaging.Status) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeRoute) {
		return nil, messaging.NewStatus(messaging.StatusInvalidContent, "")
	}
	return messaging.New[*Route](m.Content)
}

func NewExchangeMessage(ex Exchange) *messaging.Message {
	return messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent).SetContent(ContentTypeExchange, "", ex)
}

func ExchangeContent(m *messaging.Message) (Exchange, *messaging.Status) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeExchange) {
		return nil, messaging.NewStatus(messaging.StatusInvalidContent, "")
	}
	return messaging.New[Exchange](m.Content)
}
