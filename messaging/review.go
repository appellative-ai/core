package messaging

import (
	"errors"
	"sync/atomic"
	"time"
)

const (
	ContentTypeReview = "application/x-review"
	defaultDuration   = 1
)

func NewReviewMessage(review *Review) *Message {
	return NewMessage(ChannelControl, ConfigEvent).SetContent(ContentTypeReview, review)
}

func ReviewContent(m *Message) (*Review, *Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeReview) {
		return nil, NewStatus(StatusInvalidContent, errors.New("invalid content"))
	}
	return New[*Review](m.Content)
}

// Review - maybe add Task??
type Review struct {
	started  bool
	duration time.Duration
	expired  *atomic.Bool
	ticker   *Ticker
}

func NewReview(minutes int) *Review {
	if minutes <= 0 {
		minutes = defaultDuration
	}
	return newReview(time.Minute * time.Duration(minutes))
}

func newReview(dur time.Duration) *Review {
	r := new(Review)
	r.expired = new(atomic.Bool)
	r.expired.Store(true)
	r.duration = dur
	return r
}

func (r *Review) Started() bool {
	return r.started
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
