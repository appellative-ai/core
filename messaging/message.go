package messaging

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	StartupEvent  = "core:event/startup"
	ShutdownEvent = "core:event/shutdown"
	PauseEvent    = "core:event/pause"  // disable data channel receive
	ResumeEvent   = "core:event/resume" // enable data channel receive
	ConfigEvent   = "core:event/config"
	StatusEvent   = "core:event/status"

	//ObservationEvent = "event:observation"
	//TickEvent        = "event:tick"
	//DataChangeEvent  = "event:data-change"

	ChannelMaster   = "master"
	ChannelEmissary = "emissary"
	ChannelControl  = "ctrl"
	ChannelData     = "data"

	XTo        = "x-to"
	XFrom      = "x-from"
	XName      = "x-name"
	XChannel   = "x-channel"
	XRelatesTo = "x-relates-to"

	ContentType       = "Content-Type"
	ContentTypeMap    = "application/x-map"
	ContentTypeStatus = "application/x-status"
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
type Handler func(msg *Message)

// Message - message
type Message struct {
	Header  http.Header //map[string]string
	Content any
	Expiry  time.Time
	Reply   Handler
}

func NewMessage(channel, name string) *Message {
	m := new(Message)
	m.Header = make(http.Header) //map[string]string)
	m.Header.Set(XChannel, channel)
	m.Header.Set(XName, name)
	return m
}

/*
func NewMessageWithError(channel, name string, err error) *Message {
	m := NewMessage(channel, name)
	m.SetContent(ContentTypeError, err)
	return m
}
*/

func NewAddressableMessage(channel, name, to, from string) *Message {
	m := NewMessage(channel, name)
	m.Header.Add(XTo, to)
	m.Header.Add(XFrom, from)
	return m
}

func (m *Message) String() string {
	return fmt.Sprintf("[chan:%v] [from:%v] [to:%v] [%v]", m.Channel(), m.From(), m.To(), m.Name())
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

func (m *Message) Name() string {
	return m.Header.Get(XName)
}

func (m *Message) Channel() string {
	return m.Header.Get(XChannel)
}

func (m *Message) SetChannel(channel string) *Message {
	m.Header.Set(XChannel, channel)
	return m
}

func (m *Message) SetContentType(contentType string) *Message {
	if len(contentType) == 0 {
		return m //errors.New("error: content type is empty")
	}
	m.Header.Set(ContentType, contentType)
	return m
}

func (m *Message) ContentType() string {
	return m.Header.Get(ContentType)
}

func (m *Message) SetContent(contentType string, content any) error {
	if len(contentType) == 0 {
		return errors.New("error: content type is empty")
	}
	if content == nil {
		return errors.New("error: content is nil")
	}
	m.Content = content
	m.Header.Set(ContentType, contentType)
	return nil
}

func NewConfigMapMessage(cfg map[string]string) *Message {
	m := NewMessage(ChannelControl, ConfigEvent)
	m.SetContent(ContentTypeMap, cfg)
	return m
}

func ConfigMapContent(m *Message) map[string]string {
	if m.Name() != ConfigEvent || m.ContentType() != ContentTypeMap {
		return nil
	}
	if cfg, ok := m.Content.(map[string]string); ok {
		return cfg
	}
	return nil
}

func NewStatusMessage(status *Status, relatesTo string) *Message {
	m := NewMessage(ChannelControl, StatusEvent)
	m.SetContent(ContentTypeStatus, status)
	if relatesTo != "" {
		m.Header.Set(XRelatesTo, relatesTo)
	}
	return m
}

func StatusContent(m *Message) (*Status, string) {
	if m.Name() != StatusEvent || m.ContentType() != ContentTypeStatus {
		return nil, ""
	}
	if s, ok := m.Content.(*Status); ok {
		return s, m.RelatesTo()
	}
	return nil, ""
}

// Reply - function used by message recipient to reply with a Status
func Reply(msg *Message, status *Status, from string) {
	if msg == nil || status == nil || msg.Reply == nil {
		return
	}
	m := NewStatusMessage(status, msg.Name())
	m.Header.Set(XFrom, from)
	msg.Reply(m)
}
