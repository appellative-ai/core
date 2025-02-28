package messaging

// OnShutdown - add functions to be run on shutdown
type OnShutdown interface {
	Add(func())
}

func AddShutdown(curr, next func()) func() {
	if next == nil {
		return nil
	}
	if curr == nil {
		curr = next
	} else {
		// !panic
		prev := curr
		curr = func() {
			prev()
			next()
		}
	}
	return curr
}

/*
defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in agent.Shutdown() : %v\n", r)
		}
	}()

*/
