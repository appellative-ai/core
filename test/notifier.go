package test

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messagingx"
)

type Notifier interface {
	messagingx.Notifier
	Status() *aspect.Status
	Reset()
}

type statusT struct {
	status *aspect.Status
}

func (s *statusT) Notify(status *aspect.Status) *aspect.Status {
	s.status = status
	return status
}

func (s *statusT) Status() *aspect.Status {
	return s.status
}

func (s *statusT) Reset() {
	s.status = nil
}

func NewNotifier() Notifier {
	return new(statusT)
}
