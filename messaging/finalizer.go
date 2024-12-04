package messaging

import "time"

type Finalizer interface {
	IsFinalized() bool
}

func IsFinalized(attempts int, sleep time.Duration, finalized func() bool) bool {
	for i := attempts; i > 0; i-- {
		if finalized() {
			return true
		}
		time.Sleep(sleep)
	}
	return false
}
