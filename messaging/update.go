package messaging

import "github.com/appellative-ai/core/std"

func UpdateReview(name string, review **Review, m *Message) {
	if m == nil || review == nil || m.ContentType() != ContentTypeReview {
		return
	}
	review2, status := ReviewContent(m)
	if !status.OK() {
		Reply(m, status, name)
		return
	}
	*review = review2
	Reply(m, std.StatusOK, name)
}

func UpdateMap(name string, fn func(cfg map[string]string), m *Message) {
	if m == nil || fn == nil || m.ContentType() != ContentTypeMap {
		return
	}
	cfg, status := MapContent(m)
	if !status.OK() {
		Reply(m, status, name)
		return
	}
	fn(cfg)
	Reply(m, std.StatusOK, name)
}

func UpdateDispatcher(name string, d *Dispatcher, m *Message) {
	if m == nil || d == nil || m.ContentType() != ContentTypeDispatcher {
		return
	}
	dsp, ok := DispatcherContent(m)
	if !ok {
		Reply(m, std.NewStatus(std.StatusInvalidArgument, "", nil), name)
		return
	}
	*d = dsp
	Reply(m, std.StatusOK, name)
}

func UpdateAgent(name string, fn func(agent Agent), m *Message) {
	if m == nil || fn == nil || m.ContentType() != ContentTypeAgent {
		return
	}
	a, status := AgentContent(m)
	if !status.OK() {
		Reply(m, status, name)
		return

	}
	fn(a)
	Reply(m, std.StatusOK, name)
}
