package messaging

import (
	"github.com/behavioral-ai/core/fmtx"
	"sync/atomic"
	"time"
)

const (
	ContentTypeReview = "application/x-review"
	defaultDuration   = time.Minute * 1
)

func NewReviewMessage(review *Review) *Message {
	m := NewMessage(ChannelControl, ConfigEvent)
	m.SetContent(ContentTypeReview, review)
	return m
}

func ReviewContent(m *Message) *Review {
	if m.Name() != ConfigEvent || m.ContentType() != ContentTypeReview {
		return nil
	}
	if v, ok := m.Body.(*Review); ok {
		return v
	}
	return nil
}

// Review - maybe add Task??
type Review struct {
	started  bool
	duration time.Duration
	expired  *atomic.Bool
	ticker   *Ticker
}

func NewReview(dur string) *Review {
	r := new(Review)
	r.expired = new(atomic.Bool)
	r.expired.Store(true)
	if dur == "" {
		return r
	}
	d, err := fmtx.ParseDuration(dur)
	if err != nil {
		return r
	}
	r.duration = d
	if r.duration < defaultDuration {
		r.duration = defaultDuration
	}
	return r
}

func (r *Review) Started() bool {
	return r.started
}

func (r *Review) Scheduled() bool {
	return r.duration != 0
}

func (r *Review) Expired() bool {
	return r.expired.Load()
}

func (r *Review) Start() {
	r.ticker = NewTicker(ChannelControl, r.duration)
	r.expired.Store(false)
	r.started = true
	go reviewAttend(r)
}

func reviewAttend(r *Review) {
	for {
		select {
		case <-r.ticker.C():
			r.expired.Store(true)
			return
		default:
		}
	}
}
