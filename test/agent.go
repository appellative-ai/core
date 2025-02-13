package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	agentId string
	ch      *messaging.Channel
	//shutdownFunc func()
}

func NewAgent(uri string) messaging.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = messaging.NewEmissaryChannel(true)
	return a
}

func NewAgentWithChannel(uri string, ch *messaging.Channel) messaging.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agentT) Uri() string { return t.agentId }
func (t *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	t.ch.C <- m
}

func (t *agentT) IsFinalized() bool { return t.ch.IsFinalized() }

// Notify - status notifications
func (t *agentT) Notify(status error) {
	//var e aspect.Output
	fmt.Printf("test: opsAgent.Handle() -> [status:%v]\n", status)
	//status.Handled = true
	//e.Handle(status)
	return
}

// Trace - activity tracing
func (t *agentT) Trace(agent messaging.Agent, channel, event, activity string) {
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
				case messaging.ShutdownEvent:
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
	msg := messaging.NewControlMessage(t.Uri(), t.Uri(), messaging.ShutdownEvent)
	t.ch.Enable()
	t.ch.C <- msg
}

func (t *agentT) finalize() {
	t.ch.Close()
	t.ch = nil
}
