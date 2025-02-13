package test

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
)

type agentT struct {
	agentId string
	ch      *messagingx.Channel
	//shutdownFunc func()
}

func NewAgent(uri string) messagingx.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = messagingx.NewEmissaryChannel(true)
	return a
}

func NewAgentWithChannel(uri string, ch *messagingx.Channel) messagingx.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agentT) Uri() string { return t.agentId }
func (t *agentT) Message(m *messagingx.Message) {
	if m == nil {
		return
	}
	t.ch.C <- m
}

func (t *agentT) IsFinalized() bool { return t.ch.IsFinalized() }

// Notify - status notifications
func (t *agentT) Notify(status *aspect.Status) *aspect.Status {
	var e aspect.Output
	//fmt.Printf("test: opsAgent.Handle() -> [status:%v]\n", status)
	//status.Handled = true
	return e.Handle(status)
}

// Trace - activity tracing
func (t *agentT) Trace(agent messagingx.Agent, channel, event, activity string) {
	trace(agent, channel, event, activity)
}

// Add - add a shutdown function
//func (t *agentT) Add(f func()) { t.shutdownFunc = messagingx.AddShutdown(t.shutdownFunc, f) }

func (t *agentT) Run() {
	go func() {
		for {
			select {
			case msg := <-t.ch.C:
				switch msg.Event() {
				case messagingx.ShutdownEvent:
					t.finalize()
					return
				default:
				}
			default:
			}
		}
	}()
}

func (t *agentT) Shutdown() {
	//if t.shutdownFunc != nil {
	//	t.shutdownFunc()
	//}
	msg := messagingx.NewControlMessage(t.Uri(), t.Uri(), messagingx.ShutdownEvent)
	t.ch.Enable()
	t.ch.C <- msg
}

func (t *agentT) finalize() {
	t.ch.Close()
	t.ch = nil
}
