package messaging

import "fmt"

type testShutdown struct {
	agentId    string
	shutdownFn func()
}

// Enqueue - append a shutdown function to a queue
func (t *testShutdown) Enqueue(f func()) {
	t.shutdownFn = AddShutdown(t.shutdownFn, f)
}

// Push - push a shutdown function onto a stack
func (t *testShutdown) Push(f func()) {
	if t.shutdownFn == nil {
		t.shutdownFn = func() {}
	}
	t.shutdownFn = AddShutdown(f, t.shutdownFn)
}

// Shutdown - shutdown the agent
func (t *testShutdown) Shutdown() {
	if t.shutdownFn != nil {
		t.shutdownFn()
	}
}

func ExampleShutdown_Queue() {
	t := &testShutdown{agentId: "agent/test"}

	t.Enqueue(func() { fmt.Printf("test: Enqueue(1) -> func(%v)\n", 1) })
	t.Enqueue(func() { fmt.Printf("test: Enqueue(2) -> func(%v)\n", 2) })
	t.Enqueue(func() { fmt.Printf("test: Enqueue(3) -> func(%v)\n", 3) })
	t.Enqueue(func() { fmt.Printf("test: Enqueue(4) -> func(%v)\n", 4) })
	t.shutdownFn()

	//Output:
	//test: Enqueue(1) -> func(1)
	//test: Enqueue(2) -> func(2)
	//test: Enqueue(3) -> func(3)
	//test: Enqueue(4) -> func(4)

}

func ExampleShutdown_Stack() {
	t := &testShutdown{agentId: "agent/test"}

	t.Push(func() { fmt.Printf("test: Push(1) -> func(%v)\n", 1) })
	t.Push(func() { fmt.Printf("test: Push(2) -> func(%v)\n", 2) })
	t.Push(func() { fmt.Printf("test: Push(3) -> func(%v)\n", 3) })
	t.Push(func() { fmt.Printf("test: Push(4) -> func(%v)\n", 4) })
	t.shutdownFn()

	//Output:
	//test: Push(4) -> func(4)
	//test: Push(3) -> func(3)
	//test: Push(2) -> func(2)
	//test: Push(1) -> func(1)

}

/*
if sd, ok1 := t.(OnShutdown); ok1 {
	sd.S
}
	sd.Add(func() {
		d.m.Delete(m.Uri())
}


*/
