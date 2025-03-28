package access

import (
	"strconv"
	"time"
)

type Thresholds struct {
	Timeout   any
	RateLimit any
	Redirect  any
}

func (t Thresholds) timeout() time.Duration {
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

func (t Thresholds) rateLimit() float64 {
	var limit float64 = -1

	if t.RateLimit == nil {
		return limit
	}
	return limit
}

func (t Thresholds) redirect() int {
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
