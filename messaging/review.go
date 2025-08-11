package messaging

import (
	"errors"
	"github.com/appellative-ai/core/std"
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

func ReviewContent(m *Message) (*Review, *std.Status) {
	if !ValidContent(m, ConfigEvent, ContentTypeReview) {
		return nil, std.NewStatus(std.StatusInvalidContent, "", errors.New("invalid content"))
	}
	return std.New[*Review](m.Content)
}

// Review - maybe add Task??
type Review struct {
	started  bool
	duration time.Duration
	expired  atomic.Bool
	ticker   *Ticker
}

func NewReview() *Review {
	//if minutes <= 0 {
	//	minutes = defaultDuration
	//}
	r := new(Review)
	return r //newReview(time.Minute * time.Duration(minutes))
}

/*
func newReview(dur time.Duration) *Review {
	r := new(Review)
	r.expired = new(atomic.Bool)
	r.expired.Store(true)
	r.duration = dur
	return r
}


*/

func (r *Review) Started() bool {
	return r.started
}

func (r *Review) Expired() bool {
	return r.expired.Load()
}

func (r *Review) Start(dur time.Duration) {
	if r.started {
		return
	}
	if dur <= 0 {
		dur = defaultDuration
	}
	r.ticker = NewTicker(ChannelControl, dur)
	r.expired.Store(false)
	r.started = true
	r.duration = dur
	go reviewAttend(r)
}

func reviewAttend(r *Review) {
	for {
		select {
		case <-r.ticker.T.C:
			r.expired.Store(true)
			return
		default:
		}
	}
}
