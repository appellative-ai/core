package messaging

/*
func ExampleUpdateMap() {
	UpdateMap("", nil, nil)
	fmt.Printf("test: UpdateMap() -> nil update fn\n")

	fn := func(cfg map[string]string) {
		fmt.Printf("test: UpdateMap() -> %v\n", cfg)
	}
	UpdateMap("", fn, nil)
	fmt.Printf("test: UpdateMap() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateMap("", fn, m)
	fmt.Printf("test: UpdateMap() -> invalid content type\n")

	cfg := map[string]string{"name": "value"}
	m = NewMapMessage(cfg)
	UpdateMap("", fn, m)

	//Output:
	//test: UpdateMap() -> nil update fn
	//test: UpdateMap() -> nil message
	//test: UpdateMap() -> invalid content type
	//test: UpdateMap() -> map[name:value]

}

func ExampleUpdateReview() {
	UpdateReview("", nil, nil)
	fmt.Printf("test: UpdateReview() -> nil Review\n")

	review := NewReview()
	UpdateReview("", &review, nil)
	fmt.Printf("test: UpdateReview() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateReview("", &review, m)
	fmt.Printf("test: UpdateReview() -> invalid content type\n")

	before := review
	m = NewReviewMessage(NewReview())
	UpdateReview("", &review, m)
	fmt.Printf("test: UpdateReview() -> [original:%v] [updated:%v]\n", before.duration, review.duration)

	//Output:
	//test: UpdateReview() -> nil Review
	//test: UpdateReview() -> nil message
	//test: UpdateReview() -> invalid content type
	//test: UpdateReview() -> [original:5m0s] [updated:10m0s]

}

func ExampleUpdateDispatcher() {
	UpdateDispatcher("", nil, nil)
	fmt.Printf("test: UpdateDispatcher() -> nil Dispatcher\n")

	d := NewTraceDispatcher()
	UpdateDispatcher("", &d, nil)
	fmt.Printf("test: UpdateDispatcher() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateDispatcher("", &d, m)
	fmt.Printf("test: UpdateDispatcher() -> invalid content type\n")

	var d2 Dispatcher
	m = NewDispatcherMessage(NewTraceDispatcher())
	UpdateDispatcher("", &d2, m)
	fmt.Printf("test: UpdateDispatcher() -> [original:%v] [updated:%v]\n", nil, d2)

	//Output:
	//test: UpdateDispatcher() -> nil Dispatcher
	//test: UpdateDispatcher() -> nil message
	//test: UpdateDispatcher() -> invalid content type
	//test: UpdateDispatcher() -> [original:<nil>] [updated:&{true  map[]}]

}

func ExampleUpdateAgent() {
	UpdateAgent("", nil, nil)
	fmt.Printf("test: UpdateAgent() -> nil update fn\n")

	fn := func(agent Agent) {
		fmt.Printf("test: UpdateAgent() -> %v\n", agent)
	}
	UpdateAgent("", fn, nil)
	fmt.Printf("test: UpdateAgent() -> nil message\n")

	m := NewMessage(ChannelControl, "test-message")
	UpdateAgent("", fn, m)
	fmt.Printf("test: UpdateAgent() -> invalid content type\n")

	a := NewAgent("test:agent", func(m *Message) {})
	m = NewAgentMessage(a)
	UpdateAgent("", fn, m)

	//Output:
	//test: UpdateAgent() -> nil update fn
	//test: UpdateAgent() -> nil message
	//test: UpdateAgent() -> invalid content type
	//test: UpdateAgent() -> test:agent

}


*/
