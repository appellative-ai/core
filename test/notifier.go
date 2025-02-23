package test

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func Notify(e messaging.Event) {
	fmt.Printf("notify-> [event:%v] [msg:%v] [src:%v] [agent:%v]", e.Name(), e.Content(), e.Source(), e.AgentId())
	//fmt.Printf("notify-> [status:%v]\n", status)
	//return status
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
