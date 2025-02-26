package messaging

const (
	EmissaryChannel = "emissary"
	MasterChannel   = "master"
	PrimaryChannel  = "primary"
	ControlChannel  = "ctrl"
	DataChannel     = "data"
)

type Channel struct {
	name string
	C    chan *Message
}

func NewChannel(name string) *Channel {
	c := new(Channel)
	c.name = name
	c.C = make(chan *Message, ChannelSize)
	return c
}

func NewEmissaryChannel() *Channel {
	return NewChannel(EmissaryChannel)
}

func NewMasterChannel() *Channel {
	return NewChannel(MasterChannel)
}

/*
func NewPrimaryChannel(enabled bool) *Channel {
	return NewChannel(PrimaryChannel, enabled)
}
*/

func (c *Channel) String() string { return c.Name() }
func (c *Channel) Name() string   { return c.name }
func (c *Channel) IsClosed() bool { return c.C == nil }
func (c *Channel) Close() {
	if c.C != nil {
		close(c.C)
		c.C = nil
	}
}

func (c *Channel) Send(m *Message) {
	if m != nil {
		c.C <- m
	}
}

//func (c *Channel) IsEnabled() bool { return c.enabled }
//func (c *Channel) Enable()         { c.enabled = true }
//func (c *Channel) Disable()        { c.enabled = false }
