package messaging

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	exchangeFinalizeDuration = time.Second * 5
	exchangeFinalizeAttempts = 3
)

// Mailbox - mailbox functionality
/*
type Mailbox interface {
	Uri() string
	Message(m *Message)
}

*/

// Exchange - exchange directory
type Exchange struct {
	m *sync.Map
}

// NewExchange - create a new controller2
func NewExchange() *Exchange {
	e := new(Exchange)
	e.m = new(sync.Map)
	return e
}

// Count - number of agents
func (e *Exchange) Count() int {
	count := 0
	e.m.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// List - a list of agent uri's
func (e *Exchange) List() []string {
	var uri []string
	e.m.Range(func(key, value any) bool {
		if str, ok := key.(string); ok {
			uri = append(uri, str)
		}
		return true
	})
	sort.Strings(uri)
	return uri
}

// Message - send a message
func (e *Exchange) Message(msg *Message) error {
	// TO DO : authenticate shutdown control message
	//if msg != nil && msg.Name() == ShutdownEvent {
	//	return nil
	//}
	if msg == nil {
		return errors.New(fmt.Sprintf("error: exchange.Message() failed as message is nil"))
	}
	a := e.Get(msg.To())
	if a == nil {
		return errors.New(fmt.Sprintf("error: exchange.Message() failed as the message To is empty or invalid : [%v]", msg.To()))
	}
	a.Message(msg)
	return nil
}

// Broadcast - broadcast a message to all entries, deleting the entry if the message event is Shutdown
func (e *Exchange) Broadcast(msg *Message) {
	if msg == nil {
		return //errors.New(fmt.Sprintf("error: exchange.Broadcast() failed as message is nil"))
	}
	// TODO : Need to disallow shutdown message??
	//if msg.Name() == ShutdownEvent {
	//	return
	//}
	for _, uri := range e.List() {
		a := e.Get(uri)
		if a == nil {
			continue
		}
		a.Message(msg)
		//if msg.Name() == ShutdownEvent {
		//	d.m.Delete(uri)
		//}
	}
}

// registerMailbox - register a mailbox
/*
func (e *Exchange) registerMailbox(m Mailbox) error {
	if m == nil {
		return errors.New("error: exchange.Register() agent is nil")
	}
	_, ok := e.m.Load(m.Uri())
	if ok {
		return errors.New(fmt.Sprintf("error: exchange.Register() agent already exists: [%v]", m.Uri()))
	}
	e.m.Store(m.Uri(), m)
	if sd, ok1 := m.(OnShutdown); ok1 {
			sd.Add(func() {
				e.m.Delete(m.Uri())
			})
		}
	return nil
}

 // GetMailbox - find a mailbox
func (e *Exchange) GetMailbox(uri string) Mailbox {
	if len(uri) == 0 {
		return nil
	}
	v, ok1 := e.m.Load(uri)
	if !ok1 {
		return nil
	}
	if a, ok2 := v.(Mailbox); ok2 {
		return a
	}
	return nil
}

*/

// Register - register an agent
func (e *Exchange) Register(agent Agent) error {
	if agent == nil {
		return errors.New("error: exchange.Register() agent is nil")
	}
	if agent.Name() == "" {
		return errors.New("error: exchange.Register() agent Name is empty")
	}
	_, ok := e.m.Load(agent.Name())
	if ok {
		return errors.New(fmt.Sprintf("error: exchange.Register() agent already exists: [%v]", agent.Name()))
	}
	e.m.Store(agent.Name(), agent)
	/*
		if sd, ok1 := agent.(OnShutdown); ok1 {
			sd.Add(func() {
				e.m.Delete(agent.Name())
			})
		}

	*/
	return nil
}

// Get - find an agent
func (e *Exchange) Get(name string) Agent {
	if len(name) == 0 {
		return nil
	}
	v, ok1 := e.m.Load(name)
	if !ok1 {
		return nil
	}
	if a, ok2 := v.(Agent); ok2 {
		return a
	}
	return nil
}

// Shutdown - shutdown all agents
func (e *Exchange) Shutdown() {
	//go func() {
	for _, name := range e.List() {
		a := e.Get(name)
		if a == nil {
			continue
		}
		a.Message(ShutdownMessage)
		// TODO: verify
		e.m.Delete(name)
	}
	//	}()
}

// IsFinalized - determine if all agents have been shutdown and removed from the exchange
func (e *Exchange) IsFinalized() bool {
	for i := exchangeFinalizeAttempts; i > 0; i-- {
		if e.Count() == 0 {
			return true
		}
		time.Sleep(exchangeFinalizeDuration)
	}
	return false
}
