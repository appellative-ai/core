package logx

import (
	"fmt"
	"github.com/appellative-ai/core/std"
	"log"
	"net/http"
	"time"
)

//ContentTypeOperators = "application/x-log-operators"
//DefaultRoute         = "host"
//RequestIdName = "x-request-id"
//	IngressTraffic       = "ingress"

const (
	EgressTraffic = "egress"

	ThresholdName = "x-threshold"
	RateLimitName = "rate-limit"
	TimeoutName   = "timeout"
	RedirectName  = "redirect"
	CachedName    = "cached"

	failsafeUri     = "https://invalid-uri.com"
	contentEncoding = "Content-Encoding"
)

// Operator - configuration of logging entries
type Operator struct {
	Name  string
	Value string
}

// Request - request interface for non HTTP traffic
type Request interface {
	Url() string
	Header() http.Header
	Method() string
	Protocol() string
}

// Response - response interface for non HTTP traffic
type Response interface {
	StatusCode() int
	Header() http.Header
}

// LogAccess - access traffic
func LogAccess(operators []Operator, traffic string, start time.Time, duration time.Duration, route string, req any, resp any) {
	if len(operators) == 0 {
		operators = defaultOperators
	}
	e := newEvent(traffic, start, duration, route, req, resp)
	s := writeJson(operators, e)
	log.Printf("%v\n", s)
}

// LogEgress - egress traffic
func LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	var r *http.Response
	if duration > 0 {
		r = buildResponse(resp)
		SetTimeout(r.Header, timeout)
	}
	LogAccess(nil, EgressTraffic, start, duration, route, req, r)
}

// LogStatus - log status
func LogStatus(status *std.Status) {
	log.Printf("%v\n", status)
}

func SetTimeout(h http.Header, v time.Duration) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", TimeoutName, v))
}

func SetRateLimit(h http.Header, v float64) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", RateLimitName, v))
}

func SetRedirect(h http.Header, v int) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", RedirectName, v))
}

func SetCached(h http.Header, v bool) {
	if h == nil {
		return
	}
	h.Add(ThresholdName, fmt.Sprintf("%v=%v", CachedName, v))
}

func RemoveThresholds(h http.Header) {
	if h == nil {
		return
	}
	h.Del(ThresholdName)
}
