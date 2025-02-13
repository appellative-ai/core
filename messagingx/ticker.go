package messaging

import "time"

const (
	PrimaryTicker          = "PRIMARY"
	tickerFinalizeAttempts = 3
	tickerFinalizeDuration = time.Second * 5
)

type Ticker struct {
	name     string
	duration time.Duration
	original time.Duration
	ticker   *time.Ticker
}

func NewTicker(name string, duration time.Duration) *Ticker {
	t := new(Ticker)
	t.name = name
	t.duration = duration
	t.original = duration
	t.ticker = time.NewTicker(duration)
	return t
}

func NewPrimaryTicker(duration time.Duration) *Ticker {
	return NewTicker(PrimaryTicker, duration)
}

func (t *Ticker) String() string          { return t.Name() }
func (t *Ticker) Name() string            { return t.name }
func (t *Ticker) Duration() time.Duration { return t.duration }
func (t *Ticker) C() <-chan time.Time     { return t.ticker.C }
func (t *Ticker) IsFinalized() bool {
	return IsFinalized(tickerFinalizeAttempts, tickerFinalizeDuration, t.IsStopped)
}

func (t *Ticker) Start(newDuration time.Duration) {
	if newDuration <= 0 {
		newDuration = t.duration
	} else {
		t.duration = newDuration
	}
	t.ticker.Stop()
	t.ticker.Reset(newDuration)
}

func (t *Ticker) Reset() {
	t.Start(t.original)
}

func (t *Ticker) IsStopped() bool {
	return t.ticker == nil
}

func (t *Ticker) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
		t.ticker = nil
	}
}
