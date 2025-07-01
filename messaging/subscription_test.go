package messaging

import (
	"fmt"
	"time"
)

func ExampleSubscription_Valid() {
	s := NewSubscription("create-from", ChannelControl, SubscriptionCreateEvent, "")

	path := ""
	fmt.Printf("test: Valid(\"%v\") -> [sub-path:%v] [valid:%v]\n", path, s.Path, s.Valid(path))

	path = "/google/search"
	fmt.Printf("test: Valid(\"%v\") -> [sub-path:%v] [valid:%v]\n", path, s.Path, s.Valid(path))

	s = NewSubscription("create-from", ChannelControl, SubscriptionCreateEvent, "/google/search")

	path = ""
	fmt.Printf("test: Valid(\"%v\") -> [sub-path:%v] [valid:%v]\n", path, s.Path, s.Valid(path))

	path = "/google/search"
	fmt.Printf("test: Valid(\"%v\") -> [sub-path:%v] [valid:%v]\n", path, s.Path, s.Valid(path))

	//Output:
	//test: Valid("") -> [sub-path:] [valid:true]
	//test: Valid("/google/search") -> [sub-path:] [valid:true]
	//test: Valid("") -> [sub-path:/google/search] [valid:false]
	//test: Valid("/google/search") -> [sub-path:/google/search] [valid:true]

}

func ExampleSubscriptionMessage() {
	m := NewSubscriptionCreateMessage("create-to", NewSubscription("create-from", ChannelControl, SubscriptionCreateEvent, ""))
	s, status := SubscriptionCreateContent(m)
	fmt.Printf("test: NewSubscriptionCreateMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Name, s, status)

	m = NewSubscriptionCancelMessage("cancel-to", "cancel-from", SubscriptionCancelEvent)
	s, status = SubscriptionCancelContent(m)
	fmt.Printf("test: NewSubscriptionCancelMessage() -> [%v] [%v] [%v] [%v]\n", m.To(), m.Name, s, status)

	//Output:
	//test: NewSubscriptionCreateMessage() -> [create-to] [core:event/subscription-create] [{ ctrl core:event/subscription-create create-from}] [OK]
	//test: NewSubscriptionCancelMessage() -> [cancel-to] [core:event/subscription-cancel] [{  core:event/subscription-cancel cancel-from}] [OK]

}

func _ExampleCatalog_Create() {
	c := new(Catalog)

	m := NewSubscriptionCreateMessage(publisherName, Subscription{From: "", Channel: ChannelControl, Name: publishEvent})
	err := c.CreateWithMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, Subscription{From: subscriberName, Channel: ChannelControl})
	err = c.CreateWithMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, Subscription{From: subscriberName, Channel: ChannelControl, Name: publishEvent})
	err = c.CreateWithMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	m = NewSubscriptionCreateMessage(publisherName, Subscription{From: subscriberName, Channel: ChannelControl, Name: publishEvent})
	err = c.CreateWithMessage(m)
	fmt.Printf("test: Catalog() -> [err:%v]\n", err)

	//Output:
	//test: Catalog() -> [err:invalid subscription: from or Name is empty]
	//test: Catalog() -> [err:invalid subscription: from or Name is empty]
	//test: Catalog() -> [err:<nil>]
	//test: Catalog() -> [err:invalid subscription: subscription is a duplicate [subscriber] [Name:publish]]

}

func _ExampleCatalog_Lookup() {
	event1 := "event:publish-1"
	event2 := "event:test"
	c := new(Catalog)

	err := c.Create(Subscription{Name: publishEvent, From: subscriberName})
	if err != nil {
		fmt.Printf("test: Catalog() -> [err:%v]\n", err)
	}

	err = c.Create(Subscription{Name: publishEvent, From: "subscriber-1"})
	if err != nil {
		fmt.Printf("test: Catalog() -> [err:%v]\n", err)
	}

	err = c.Create(Subscription{Name: event1, From: subscriberName})
	if err != nil {
		fmt.Printf("test: Catalog() -> [err:%v]\n", err)
	}

	subs, ok := c.Lookup(event2)
	fmt.Printf("test: Catalog(\"%v\") -> [subs:%v] [ok:%v]\n", event2, subs, ok)

	subs, ok = c.Lookup(event1)
	fmt.Printf("test: Catalog(\"%v\") -> [subs:%v] [ok:%v]\n", event1, subs, ok)

	subs, ok = c.Lookup(publishEvent)
	fmt.Printf("test: Catalog(\"%v\") -> [subs:%v] [ok:%v]\n", publishEvent, subs, ok)

	//Output:
	//test: Catalog("Name:test") -> [subs:[]] [ok:false]
	//test: Catalog("Name:publish-1") -> [subs:[{Name:publish-1 subscriber}]] [ok:true]
	//test: Catalog("Name:publish") -> [subs:[{Name:publish subscriber} {Name:publish subscriber-1}]] [ok:true]

}

func ExampleCatalog_Cancel_1() {
	c := new(Catalog)

	// create 1 subscription and cancel
	m := NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, publishEvent, ""))
	err := c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	c.CancelWithMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	//Output:
	//test: Catalog.PreCancel()  -> [count:1]
	//test: Catalog.PostCancel() -> [count:0]

}

func ExampleCatalog_Cancel_2() {
	event1 := "Name:publish-1"
	c := new(Catalog)

	// create 2 subscriptions and cancel
	m := NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, publishEvent, ""))
	err := c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, event1, ""))
	err = c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event1)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelWithMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelWithMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	//Output:
	//test: Catalog.PreCancel()  -> [count:2]
	//test: Catalog.PostCancel() -> [count:1]
	//test: Catalog.PreCancel()  -> [count:1]
	//test: Catalog.PostCancel() -> [count:0]

}

func ExampleCatalog_Cancel_3() {
	event1 := "Name:publish-1"
	event2 := "Name:publish-2"
	c := new(Catalog)

	// create 3 subscriptions
	m := NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, publishEvent, ""))
	err := c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, event1, ""))
	err = c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}
	m = NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, event2, ""))
	err = c.CreateWithMessage(m)
	if err != nil {
		fmt.Printf("test: Catalog.Create() -> [err:%v]\n", err)
	}

	// cancel middle, first, last
	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event1)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelWithMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, publishEvent)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelWithMessage(m)
	fmt.Printf("test: Catalog.PostCancel() -> [count:%v]\n", len(c.subs))

	m = NewSubscriptionCancelMessage(publisherName, subscriberName, event2)
	fmt.Printf("test: Catalog.PreCancel()  -> [count:%v]\n", len(c.subs))
	c.CancelWithMessage(m)
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
	subscriberName  = "subscriber"
	publisherName   = "publisher"
	publishEvent    = "core:publish"
	workEvent       = "core:work"
	contentTypeItem = "content-type/x-item"
)

var (
	exchange = NewExchange()
)

type workItem struct {
	statusCode int
	duration   time.Duration
}

func newWorkItemMessage(w workItem) *Message {
	m := NewMessage(ChannelControl, workEvent)
	m.SetContent(contentTypeItem, w)
	return m
}

func workItemContent(m *Message) (workItem, bool) {
	if m.Name != workEvent || m.Content == nil || m.ContentType() != contentTypeItem {
		return workItem{}, false
	}
	if v, ok := m.Content.Value.(workItem); ok {
		return v, true
	}
	return workItem{}, false
}

type subscriber struct {
	running  bool
	emissary *Channel
}

func newSubscriber() Agent {
	s := new(subscriber)
	s.emissary = NewChannel(ChannelEmissary, ChannelSize)
	return s
}
func (s *subscriber) Name() string { return subscriberName }
func (s *subscriber) Message(m *Message) {
	if !s.running {
		if m.Name == StartupEvent {
			go s.run()
			s.running = true
		}
		return
	}
	if m.Name == ShutdownEvent {
		s.running = false
	}
	s.emissary.C <- m
}

func (s *subscriber) run() {
	for {
		select {
		case m := <-s.emissary.C:
			switch m.Name {
			case publishEvent:
				exchange.Message(NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, workEvent, "")))
				fmt.Printf("test: subscriber() -> [create] [%v]\n", workEvent)
			case workEvent:
				if work, ok := workItemContent(m); ok {
					fmt.Printf("test: subscriber() -> [received] [status-code:%v] [duration:%v]\n", work.statusCode, work.duration)
				}
			case ShutdownEvent:
				exchange.Message(NewSubscriptionCancelMessage(publisherName, subscriberName, workEvent))
				fmt.Printf("test: subscriber() -> [cancel] [%v]\n", workEvent)
				s.shutdown()
				return
			default:
				fmt.Printf("test: subscriber() -> [%v]\n", m)
			}
		default:
		}
	}
}

func (s *subscriber) shutdown() {
	s.running = false
	close(s.emissary.C)
}

type publisher struct {
	running  bool
	catalog  *Catalog
	emissary *Channel
}

func newPublisher() Agent {
	s := new(publisher)
	s.catalog = new(Catalog)
	s.emissary = NewChannel(ChannelEmissary, ChannelSize)
	return s
}
func (p *publisher) Name() string { return publisherName }
func (p *publisher) Message(m *Message) {
	if m.Name == StartupEvent && !p.running {
		p.running = true
		go p.run()
		return
	}
	if !p.running {
		fmt.Printf("test: publisher() [message:%v]\n", m.Name)
		return
	}
	p.emissary.C <- m
}
func (p *publisher) run() {
	for {
		select {
		case m := <-p.emissary.C:
			switch m.Name {
			case workEvent:
				fmt.Printf("test: publisher() -> [received] [%v]\n", m.Name)
				if subs, ok := p.catalog.Lookup(m.Name); ok {
					for _, item := range subs {
						m.AddTo(item.From)
						fmt.Printf("test: publisher() -> [published] [%v] [subscriber:%v] \n", item.Name, item.From)
						exchange.Message(m)
					}
				}
			case SubscriptionCreateEvent:
				err := p.catalog.CreateWithMessage(m)
				if err != nil {
					fmt.Printf("test: publisher() -> [err:%v]\n", err)
				} else {
					fmt.Printf("test: publisher() -> [created] [%v]\n", m.Name)
				}
			case SubscriptionCancelEvent:
				p.catalog.CancelWithMessage(m)
				fmt.Printf("test: publisher() -> [canceled] [%v]\n", m.Name)
			case ShutdownEvent:
				fmt.Printf("test: publisher() -> [%v]\n", m.Name)
				p.shutdown()
				return
			default:
				fmt.Printf("test: publisher() -> [%v]\n", m)
			}
		default:
		}
	}
}

func (p *publisher) shutdown() {
	p.running = false
	close(p.emissary.C)
}

func _ExampleSubscription_Publisher() {
	p := newPublisher()
	p.Message(StartupMessage)

	p.Message(NewSubscriptionCreateMessage(publisherName, NewSubscription(subscriberName, ChannelControl, workEvent, "")))
	time.Sleep(time.Second * 2)

	p.Message(newWorkItemMessage(workItem{statusCode: 200, duration: time.Millisecond * 1500}))
	time.Sleep(time.Second * 2)

	p.Message(NewSubscriptionCancelMessage(publisherName, subscriberName, workEvent))
	time.Sleep(time.Second * 2)

	p.Message(ShutdownMessage)

	//Output:
	//test: publisher() -> [created] [event:subscription-create]
	//test: publisher() -> [received] [event:work]
	//test: publisher() -> [published] [event:work] [subscriber:subscriber]
	//test: publisher() -> [canceled] [event:subscription-cancel]

}

func _ExampleSubscription_Subscriber() {
	s := newSubscriber()
	s.Message(StartupMessage)

	s.Message(NewMessage(ChannelEmissary, publishEvent))
	time.Sleep(time.Second * 2)

	s.Message(newWorkItemMessage(workItem{statusCode: 200, duration: time.Millisecond * 1500}))
	time.Sleep(time.Second * 2)

	s.Message(ShutdownMessage)
	time.Sleep(time.Second * 2)

	//Output:
	//test: subscriber() -> [create] [event:work]
	//test: subscriber() -> [received] [status-code:200] [duration:1.5s]
	//test: subscriber() -> [cancel] [event:work]

}

func _ExampleSubscription() {
	s := newSubscriber()
	p := newPublisher()
	exchange.Register(s)
	exchange.Register(p)
	exchange.Broadcast(StartupMessage)

	// send workItem to publisher, not sent to subscriber
	p.Message(newWorkItemMessage(workItem{statusCode: 200, duration: time.Millisecond * 1500}))
	time.Sleep(time.Second * 2)

	// subscriber create subscription
	s.Message(NewMessage(ChannelEmissary, publishEvent))
	time.Sleep(time.Second * 2)

	p.Message(newWorkItemMessage(workItem{statusCode: 200, duration: time.Millisecond * 1500}))
	time.Sleep(time.Second * 2)

	exchange.Broadcast(ShutdownMessage)
	time.Sleep(time.Second * 8)

	//Output:
	//test: publisher() -> [received] [event:work]
	//test: subscriber() -> [create] [event:work]
	//test: publisher() -> [created] [event:subscription-create]
	//test: publisher() -> [received] [event:work]
	//test: publisher() -> [published] [event:work] [subscriber:subscriber]
	//test: subscriber() -> [received] [status-code:200] [duration:1.5s]
	//test: publisher() -> [event:shutdown]
	//test: publisher() [message:event:subscription-cancel]
	//test: subscriber() -> [cancel] [event:work]

}
