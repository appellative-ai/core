package messaging

type NotifyFunc func(status *Status) *Status

type Notifier interface {
	Notify(status *Status)
}

func Notify(notifier NotifyFunc, status *Status) *Status {
	if notifier != nil {
		return notifier(status)
	}
	return status
}

/*

// OnShutdown - add functions to be run on shutdown
type OnShutdown interface {
	Add(func())
}
type test struct {
   shutdownFn func()
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



if sd, ok1 := m.(OnShutdown); ok1 {
sd.Add(func() {
d.m.Delete(m.Uri())
})
}
m.unregister = func() {
		d.m.Delete(m.uri)
	}


*/
