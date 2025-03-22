package access

import (
	"time"
)

// InternalTraffic = "internal"

const (
	EgressTraffic  = "egress"
	IngressTraffic = "ingress"

	failsafeUri     = "https://invalid-uri.com"
	XRequestId      = "x-request-id"
	XRateLimit      = "x-rate-limit"
	XRateBurst      = "x-rate-burst"
	XRedirect       = "x-redirect"
	ContentEncoding = "Content-Encoding"

	ControllerTimeout   = "TO" // Controller struct code
	ControllerRateLimit = "RL" // Controller struct code
	ControllerRedirect  = "RD" // Routing struct code
)

var (
	origin = Origin{}
	logger = defaultLog
)

// SetOrigin - initialize the origin
func SetOrigin(o Origin) {
	origin = o
}

// LogFn - log function
type LogFn func(o Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, controller Controller)

// SetLogFn - override logging
func SetLogFn(fn LogFn) {
	if fn != nil {
		logger = fn
	}
}

// RequestConstraints - Request constraints
//type RequestConstraints interface {
//	*httpx.Request | Request
//}

// ResponseConstraints - Response constraints
//type ResponseConstraints interface {
//	*httpx.Response | *aspect.Status | int
//}

// Log - access logging.
// Header.Get(XRequestId)),
// Header.Get(XRelatesTo)),
// Header.Get(LocationHeader)
func Log(traffic string, start time.Time, duration time.Duration, req any, resp any, controller Controller) {
	if logger == nil {
		return
	}
	logger(origin, traffic, start, duration, req, resp, controller)
}

/*
// FormatFunc - formatting
type FormatFunc func(o aspect.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) string

// SetFormatFunc - override formatting
func SetFormatFunc(fn FormatFunc) {
	if fn != nil {
		formatter = fn
	}
}
func DisableLogging(v bool) {
	disabled = v
}
origin    = aspect.Origin{}
	//formatter = DefaultFormat
	logger    = defaultLog
	disabled  = false
*/
