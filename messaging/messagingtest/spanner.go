package messagingtest

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type testSpanner struct {
	maxSpan  bool
	min, max time.Duration
}

func NewTestSpanner(min, max time.Duration) messaging.Spanner {
	s := new(testSpanner)
	s.min = min
	s.max = max
	return s
}

func (t *testSpanner) Span() time.Duration {
	if t.maxSpan {
		t.maxSpan = false
		return t.max
	}
	t.maxSpan = true
	return t.min
}
