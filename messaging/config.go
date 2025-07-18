package messaging

func NewMapMessage(m map[string]string) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeMap, m)
}

func MapContent(m *Message) (map[string]string, *Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeMap) {
		return nil, NewStatus(StatusInvalidContent, "")
	}
	return New[map[string]string](m.Content)
}

func NewStatusMessage(status *Status, relatesTo string) *Message {
	m := NewMessage(ChannelControl, StatusEvent).SetContent(ContentTypeStatus, status)
	if relatesTo != "" {
		m.SetRelatesTo(relatesTo)
	}
	return m
}

func StatusContent(m *Message) (*Status, string, *Status) {
	if !ValidContent(m, StatusEvent, ContentTypeStatus) {
		return nil, "", NewStatus(StatusInvalidContent, "")
	}
	t, status := New[*Status](m.Content)
	if status.OK() {
		return t, m.RelatesTo(), status
	}
	return nil, "", status
}

func NewAgentMessage(a Agent) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeAgent, a)
}

func AgentContent(m *Message) (Agent, *Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeAgent) {
		return nil, NewStatus(StatusInvalidContent, "")
	}
	return New[Agent](m.Content)
}
