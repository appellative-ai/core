package messaging

import "github.com/behavioral-ai/core/core"

type Notifier interface {
	Notify(status *core.Status) *core.Status
}

var (
	LogErrorNotifier    = new(logError)
	OutputErrorNotifier = new(outputError)
)

type logError struct{}

func (l *logError) Notify(status *core.Status) *core.Status {
	var h core.Log
	return h.Handle(status)
}

type outputError struct{}

func (o *outputError) Notify(status *core.Status) *core.Status {
	var h core.Output
	return h.Handle(status)
}
