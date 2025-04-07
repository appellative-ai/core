package messaging

import "fmt"

func ExampleSubscriptionMessage() {
	m := NewSubscriptionCreateMessage("create-to", "create-from", SubscriptionCreateEvent)
	s, ok := SubscriptionCreateContent(m)
	fmt.Printf("test: NewSubscriptionCreateMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Event(), s, ok)

	m = NewSubscriptionCancelMessage("cancel-to", "cancel-from", SubscriptionCancelEvent)
	s, ok = SubscriptionCancelContent(m)
	fmt.Printf("test: NewSubscriptionCancelMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Event(), s, ok)

	//Output:
	//test: NewSubscriptionCreateMessage() -> [create-to] [event:subscription-create] [{event:subscription-create create-from}] [true]
	//test: NewSubscriptionCancelMessage() -> [cancel-to] [event:subscription-cancel] [{event:subscription-cancel cancel-from}] [true]

}

func ExampleCatalog_Create() {
	c := new(Catalog)

	m := NewSubscriptionCreateMessage(publisherName, "", publishEvent)
	err := c.CreateMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, subscriberName, "")
	err = c.CreateMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent)
	err = c.CreateMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent)
	err = c.CreateMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	//Output:
	//test: Catalog() -> [err:invalid subscription: from or event is empty]
	//test: Catalog() -> [err:invalid subscription: from or event is empty]
	//test: Catalog() -> [err:<nil>]
	//test: Catalog() -> [err:invalid subscription: subscription is a duplicate [subscriber] [event:publish]]

}

func ExampleCatalog_Cancel_1() {
	c := new(Catalog)

	// create 1 subscription and cancel
	m := NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent)
	err := c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	//Output:
	//test: Catalog.PreCancel()  -> [count:1]
	//test: Catalog.PostCancel() -> [count:0]

}

func ExampleCatalog_Cancel_2() {
	event1 := "event:publish-1"
	c := new(Catalog)

	// create 2 subscriptions and cancel
	m := NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent)
	err := c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, subscriberName, event1)
	err = c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event1)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	//Output:
	//test: Catalog.PreCancel()  -> [count:2]
	//test: Catalog.PostCancel() -> [count:1]
	//test: Catalog.PreCancel()  -> [count:1]
	//test: Catalog.PostCancel() -> [count:0]

}

func ExampleCatalog_Cancel_3() {
	event1 := "event:publish-1"
	event2 := "event:publish-2"
	c := new(Catalog)

	// create 3 subscriptions
	m := NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent)
	err := c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, subscriberName, event1)
	err = c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, subscriberName, event2)
	err = c.CreateMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}

	// cancel middle, first, last
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event1)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event2)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	//Output:
	//test: Catalog.PreCancel()  -> [count:3]
	//test: Catalog.PostCancel() -> [count:2]
	//test: Catalog.PreCancel()  -> [count:2]
	//test: Catalog.PostCancel() -> [count:1]
	//test: Catalog.PreCancel()  -> [count:1]
	//test: Catalog.PostCancel() -> [count:0]

}

const (
	subscriberName = "subscriber"
	publisherName  = "publisher"
	publishEvent   = "event:publish"
)

var (
	exchange = NewExchange()
)

type subscriber struct {
	emissary *Channel
}

func newSubscriber() Agent {
	s := new(subscriber)
	s.emissary = NewChannel(Emissary)
	return s
}
func (s *subscriber) Uri() string { return subscriberName }
func (s *subscriber) Message(m *Message) {
	if m.Event() == StartupEvent {
		go s.run()
		exchange.Send(NewSubscriptionCreateMessage(publisherName, subscriberName, publishEvent))
		return
	}
	s.emissary.C <- m
}
func (s *subscriber) run() {
	for {
		select {
		case m := <-s.emissary.C:
			switch m.Event() {
			case ShutdownEvent:
				exchange.Send(NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent))
				s.emissary.Close()
				return
			default:
				fmt.Printf("test: Subscriber() -> [%v]\n", m)
			}
		default:
		}
	}
}

type publisher struct {
	catalog  *Catalog
	emissary *Channel
}

func newPublisher() Agent {
	s := new(publisher)
	s.catalog = new(Catalog)
	s.emissary = NewChannel(Emissary)
	return s
}
func (p *publisher) Uri() string { return publisherName }
func (p *publisher) Message(m *Message) {
	if m.Event() == StartupEvent {
		go p.run()
		return
	}
	p.emissary.C <- m
}
func (p *publisher) run() {
	for {
		select {
		case m := <-p.emissary.C:
			switch m.Event() {
			case SubscriptionCreateEvent:
				err := p.catalog.CreateMessage(m)
				if err != nil {
					fmt.Printf("test: publisher() -> [err:%v]\n", err)
				}
			case SubscriptionCancelEvent:
				p.catalog.CancelMessage(m)
			case ShutdownEvent:
				p.emissary.Close()
				return
			default:
				fmt.Printf("test: publisher() -> [%v]\n", m)
			}
		default:
		}
	}
}

func _ExampleSubscription() {

}
