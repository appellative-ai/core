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
