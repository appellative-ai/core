package messaging

func ExampleActivity() {
	a := ActivityItem{
		Agent:   nil,
		Event:   "event",
		Source:  "source",
		Content: nil,
	}
	Activity(a)

	//fmt.Printf("test: Activity() -> [%v]\n",Activity(a))

	//Output:
	//active-> 2025-03-17T19:29:39.162Z [<nil>] [event] [source] [<nil>]

}
