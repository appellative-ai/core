package rest

import "github.com/behavioral-ai/core/messaging"

const (
	ContentTypeRoute = "application/x-route"
)

func NewRouteMessage(name, uri string, ex Exchange) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	r := NewRoute(name, uri, ex)
	m.SetContent(ContentTypeRoute, "", r)
	return m
}

func RouteContent(m *messaging.Message) (*Route, bool) {
	if messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeRoute) {
		return nil, false
	}
	if v, ok := m.Content.Value.(*Route); ok {
		return v, true
	}
	return nil, false
}
