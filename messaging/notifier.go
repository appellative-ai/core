package messaging

import "github.com/behavioral-ai/core/aspect"

type Notifier interface {
	Notify(status *aspect.Status) *aspect.Status
}

var (
	LogErrorNotifier    = new(logError)
	OutputErrorNotifier = new(outputError)
)

type logError struct{}

func (l *logError) Notify(status *aspect.Status) *aspect.Status {
	var h aspect.Log
	return h.Handle(status)
}

type outputError struct{}

func (o *outputError) Notify(status *aspect.Status) *aspect.Status {
	var h aspect.Output
	return h.Handle(status)
}
