package messaging

import "fmt"

func do1HandlerFn(next Handler) Handler {
	return func(m *Message) {
		fmt.Printf("test: Do1-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do2HandlerFn(next Handler) Handler {
	return func(m *Message) {
		fmt.Printf("test: Do2-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do3HandlerFn(next Handler) Handler {
	return func(m *Message) {
		fmt.Printf("test: Do3-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

func do4HandlerFn(next Handler) Handler {
	return func(m *Message) {
		fmt.Printf("test: Do4-Handler() -> receive\n")
		if next != nil {
			next(m)
		}
	}
}

type do1Handler struct{}

func (d do1Handler) Link(next Handler) Handler {
	return do1HandlerFn(next)
}

type do2Handler struct{}

func (d do2Handler) Link(next Handler) Handler {
	return do2HandlerFn(next)
}

type do3Handler struct{}

func (d do3Handler) Link(next Handler) Handler {
	return do3HandlerFn(next)
}

type do4Handler struct{}

func (d do4Handler) Link(next Handler) Handler {
	return do4HandlerFn(next)
}

func ExampleBuildHandlerChain() {
	rec := BuildMessagingChain([]any{do1HandlerFn, do2HandlerFn, do3HandlerFn, do4HandlerFn})
	rec(ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Func() {
	rec := BuildChain[Handler, Chainable[Handler]]([]any{do1HandlerFn, do2HandlerFn, do3HandlerFn, do4HandlerFn})
	rec(ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Chainable() {
	rec := BuildChain[Handler, Chainable[Handler]]([]any{do1Handler{}, do2Handler{}, do3Handler{}, do4Handler{}})
	rec(ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive

}

func ExampleBuildChainHandler_Any() {
	rec := BuildChain[Handler, Chainable[Handler]]([]any{do1Handler{}, do2HandlerFn, do3Handler{}, do4HandlerFn})
	rec(ShutdownMessage)

	//Output:
	//test: Do1-Handler() -> receive
	//test: Do2-Handler() -> receive
	//test: Do3-Handler() -> receive
	//test: Do4-Handler() -> receive
}

func ExampleBuildChain_Invalid_Link() {
	// This will panic
	rec := BuildChain[Handler, Chainable[Handler]]([]any{"test string"})
	fmt.Printf("test: BuildChain_Empty() -> %v\n", rec)

	//Output:
	//test: BuildChain_Empty() -> <nil>

}

func _ExampleBuildChain_Empty() {
	// This will panic
	rec := BuildChain[Handler, Chainable[Handler]](nil)
	fmt.Printf("test: BuildChain_Empty() -> %v\n", rec)

	//Output:
	//test: BuildChain_Empty() -> <nil>

}
