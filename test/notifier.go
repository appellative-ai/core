package test

import (
	"github.com/behavioral-ai/core/messaging"
)

type Notifier interface {
	messaging.Notifier
	Error() error
	Reset()
}

type statusT struct {
	status error
}

func (s *statusT) Notify(err error) {
	s.status = err
}

func (s *statusT) Error() error {
	return s.status
}

func (s *statusT) Reset() {
	s.status = nil
}

func NewNotifier() Notifier {
	return new(statusT)
}
