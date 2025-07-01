package messaging

import "time"

type Ticker2 struct {
	Name string
	//duration time.Duration
	//original time.Duration
	T *time.Ticker
}

func NewTicker2(name string, duration time.Duration) *Ticker2 {
	t := new(Ticker2)
	t.Name = name
	//t.duration = duration
	//t.original = duration
	t.T = time.NewTicker(duration)
	return t
}

func (t *Ticker2) String() string { return t.Name }
