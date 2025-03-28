package access

import (
	"strconv"
	"time"
)

type Threshold struct {
	Timeout   any
	RateLimit any
	Redirect  any
}

func (t Threshold) timeout() time.Duration {
	var dur time.Duration

	if t.Timeout == nil {
		return -1
	}
	if s, ok := t.Timeout.(string); ok {
		i, _ := strconv.Atoi(s)
		dur = time.Duration(i)
	} else {
		if d, ok1 := t.Timeout.(time.Duration); ok1 {
			dur = d
		}
	}
	return dur
}

func (t Threshold) rateLimit() float64 {
	var limit float64 = -1

	if t.RateLimit == nil {
		return limit
	}
	if s, ok := t.RateLimit.(string); ok {
		i, _ := strconv.Atoi(s)
		return float64(i)
	}
	if l, ok1 := t.RateLimit.(float64); ok1 {
		return l
	}
	if l, ok1 := t.RateLimit.(int); ok1 {
		return float64(l)
	}
	return limit
}

func (t Threshold) redirect() int {
	pct := -1
	if t.Redirect == nil {
		return pct
	}
	if s, ok := t.Redirect.(string); ok {
		i, _ := strconv.Atoi(s)
		pct = i
	} else {
		if d, ok1 := t.Redirect.(int); ok1 {
			pct = d
		}
	}
	return pct
}
