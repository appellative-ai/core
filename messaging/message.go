package messaging

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	CommonCollective = "common"
	CoreDomain       = "core"
	StartupEvent     = "common:core:event/startup"
	ShutdownEvent    = "common:core:event/shutdown"
	PauseEvent       = "common:core:event/pause"  // disable data channel receive
	ResumeEvent      = "common:core:event/resume" // enable data channel receive
	ConfigEvent      = "common:core:event/config"
	StatusEvent      = "common:core:event/status"

	//ObservationEvent = "event:observation"
	//TickEvent        = "event:tick"
	//DataChangeEvent  = "event:data-change"
	//ContentType       = "Content-Type"

	ChannelMaster   = "master"
	ChannelEmissary = "emissary"
	ChannelControl  = "ctrl"
	ChannelData     = "data"

	XTo          = "x-to"
	XFrom        = "x-from"
	XChannel     = "x-channel"
	XRelatesTo   = "x-relates-to"
	XMessageName = "x-message-name" // Used in request header to reference a message

	ContentTypeMap      = "application/x-map"
	ContentTypeStatus   = "application/x-status"
	ContentTypeAgent    = "application/x-agent"
	ContentTypeTextHtml = "text/html"
	ContentTypeText     = "text/plain charset=utf-8"
	ContentTypeBinary   = "application/octet-stream"
	ContentTypeJson     = "application/json"
)

var (
	StartupMessage  = NewMessage(ChannelControl, StartupEvent)
	ShutdownMessage = NewMessage(ChannelControl, ShutdownEvent)
	PauseMessage    = NewMessage(ChannelControl, PauseEvent)
	ResumeMessage   = NewMessage(ChannelControl, ResumeEvent)

	EmissaryShutdownMessage = NewMessage(ChannelEmissary, ShutdownEvent)
	MasterShutdownMessage   = NewMessage(ChannelMaster, ShutdownEvent)
)

// Handler - uniform interface for message handling
type Handler func(*Message)

// HandleFunc - func for message handling
//type HandleFunc func(msg *Message)

// Message - message
type Message struct {
	Name    string
	Header  http.Header
	Content *Content
	Expiry  time.Time
	Reply   Handler
}

func NewMessage(channel, name string) *Message {
	m := new(Message)
	m.Name = name
	m.Header = make(http.Header)
	m.Header.Set(XChannel, channel)
	return m
}

func NewAddressableMessage(channel, name, to, from string) *Message {
	m := NewMessage(channel, name)
	m.Header.Add(XTo, to)
	m.Header.Add(XFrom, from)
	return m
}

func (m *Message) String() string {
	return fmt.Sprintf("[chan:%v] [from:%v] [to:%v] [%v]", m.Channel(), m.From(), m.To(), m.Name)
	//return fmt.Sprintf("[chan:%v] [%v]", m.Channel(), m.Name())
}

func (m *Message) RelatesTo() string {
	return m.Header.Get(XRelatesTo)
}

func (m *Message) SetRelatesTo(s string) *Message {
	m.Header.Add(XRelatesTo, s)
	return m
}

func (m *Message) To() string {
	return m.Header.Get(XTo)
}

func (m *Message) SetTo(name string) *Message {
	m.Header.Set(XTo, name)
	return m
}

func (m *Message) From() string {
	return m.Header.Get(XFrom)
}

func (m *Message) SetFrom(name string) *Message {
	m.Header.Set(XFrom, name)
	return m
}

func (m *Message) Channel() string {
	return m.Header.Get(XChannel)
}

func (m *Message) SetChannel(channel string) *Message {
	m.Header.Set(XChannel, channel)
	return m
}

/*
func (m *Message) SetContentType(contentType string) *Message {
	if len(contentType) == 0 {
		return m //errors.New("error: content type is empty")
	}
	m.Header.Set(ContentType, contentType)
	return m
}
*/

func (m *Message) ContentType() string {
	if m.Content != nil {
		return m.Content.Type
	}
	return ""
}

func (m *Message) SetContent(contentType string, content any) *Message {
	//if len(contentType) == 0 {
	//	return errors.New("error: content type is empty")
	//}
	//if content == nil {
	//	return errors.New("error: content is nil")
	//}
	m.Content = &Content{Type: contentType, Value: content}
	return m
}

func ValidContent(m *Message, name, ct string) bool {
	if m == nil || m.Name != name {
		return false
	}
	if m.Content == nil || !m.Content.Valid(ct) {
		return false
	}
	return true
}

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
		m.Header.Set(XRelatesTo, relatesTo)
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

// Reply - function used by message recipient to reply with a Status
func Reply(msg *Message, status *Status, from string) {
	if msg == nil || status == nil || msg.Reply == nil {
		return
	}
	m := NewStatusMessage(status, msg.Name)
	m.Header.Set(XFrom, from)
	msg.Reply(m)
}

// SetReply - set a message reply, using the following constraint:
//
//	type ReplyConstraints interface {
//	    Agent | HandlerNotifiable
//	}
func SetReply(msg *Message, t any) {
	if t == nil {
		msg.Reply = func(msg *Message) {
			fmt.Printf("error: generic type is nil on call to messaging.SetReply\n")
		}
		return
	}
	if fn, ok := t.(func(m *Message)); ok {
		msg.Reply = fn
		return
	}
	if agent, ok := t.(Agent); ok {
		msg.Reply = func(m *Message) {
			agent.Message(m)
		}
		return
	}
	msg.Reply = func(msg *Message) {
		fmt.Printf(fmt.Sprintf("error: generic type: %v, is invalid for messaging.SetReply\n", reflect.TypeOf(t)))
	}
}

/*
func MarshalMessage[T any](msg *Message) (t any, status *Status) {
	if msg == nil {
		return t, NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: message is nil")))
	}
	return Marshal[T](msg.Content)
}

func UnmarshalMessage[T any](msg *Message) (t any, status *Status) {
	if msg == nil {
		return t, NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: message is nil")))
	}
	return Unmarshal[T](msg.Content)
}


*/
