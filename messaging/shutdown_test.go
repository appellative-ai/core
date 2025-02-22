package messaging

type testShutdown struct {
	running    bool
	agentId    string
	shutdownFn func()
}

// Message - send message
func (t *testShutdown) Message(msg *Message) {
}

// Add - add a shutdown function
func (t *testShutdown) Add(f func()) {
	t.shutdownFn = AddShutdown(t.shutdownFn, f)
}

// Shutdown - shutdown the agent
func (t *testShutdown) Shutdown() {
	if !t.running {
		return
	}
	t.running = false
	if t.shutdownFn != nil {
		t.shutdownFn()
	}
	t.Message(NewControlMessage(t.agentId, t.agentId, ShutdownEvent))
}

func ExampleShutdown() {
	//if sd, ok1 := m.(OnShutdown); ok1 {
	//	sd.Add(func() {
	//		d.m.Delete(m.Uri())
	//	})
	//}
}
