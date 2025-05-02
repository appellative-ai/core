package eventing

func ExampleActivity() {
	a := ActivityEvent{
		Agent:   nil,
		Event:   "eventing",
		Source:  "source",
		Content: nil,
	}
	OutputActivity(a)

	//Output:
	//active-> 2025-03-17T19:29:39.162Z [<nil>] [eventing] [source] [<nil>]

}
