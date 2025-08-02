package messaging

import "github.com/appellative-ai/core/std"

func NewConfigMessage(v any) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeAny, v)
}

func ConfigContent[T any](m *Message) (t T, ok bool) {
	if m == nil || m.Content == nil || m.ContentType() != ContentTypeAny {
		return
	}
	t, ok = m.Content.Value.(T)
	return
}

func UpdateContent[T any](m *Message, t *T) bool {
	if m == nil || m.Content == nil || m.ContentType() != ContentTypeAny {
		return false
	}
	if t1, ok := m.Content.Value.(T); ok {
		*t = t1
		return true
	}
	return false
}

func NewMapMessage(m map[string]string) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeMap, m)
}

func MapContent(m *Message) (map[string]string, *std.Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeMap) {
		return nil, std.NewStatus(std.StatusInvalidContent, "", nil)
	}
	return std.New[map[string]string](m.Content)
}

func NewStatusMessage(status *std.Status, relatesTo string) *Message {
	m := NewMessage(ChannelControl, StatusEvent).SetContent(ContentTypeStatus, status)
	if relatesTo != "" {
		m.SetRelatesTo(relatesTo)
	}
	return m
}

func StatusContent(m *Message) (*std.Status, string, *std.Status) {
	if !ValidContent(m, StatusEvent, ContentTypeStatus) {
		return nil, "", std.NewStatus(std.StatusInvalidContent, "", nil)
	}
	t, status := std.New[*std.Status](m.Content)
	if status.OK() {
		return t, m.RelatesTo(), status
	}
	return nil, "", status
}

func NewAgentMessage(a Agent) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeAgent, a)
}

func AgentContent(m *Message) (Agent, *std.Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeAgent) {
		return nil, std.NewStatus(std.StatusInvalidContent, "", nil)
	}
	return std.New[Agent](m.Content)
}
