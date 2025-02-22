package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func Notify(status *messaging.Status) *messaging.Status {
	fmt.Printf("test: -> [status:%v]\n", status)
	return status
}

type Notifier interface {
	messaging.Notifier
	Error() error
	Reset()
}

type statusT struct {
	status error
}

func (s *statusT) Notify(status *messaging.Status) {
	s.status = status.Err
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
