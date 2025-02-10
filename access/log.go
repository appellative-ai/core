package access

import (
	"time"
)

const (
	InternalTraffic = "internal"
	EgressTraffic   = "egress"
	IngressTraffic  = "ingress"

	failsafeUri     = "https://invalid-uri.com"
	XRequestId      = "x-request-id"
	XRelatesTo      = "x-relates-to"
	ContentEncoding = "Content-Encoding"
	LocationHeader  = "Location"

	Primary             = "primary"
	Secondary           = "secondary"
	ControllerTimeout   = "TO" // Controller struct code
	ControllerRateLimit = "RL" // Controller struct code
	RoutingFailover     = "FO" // Routing struct code
	RoutingRedirect     = "RD" // Routing struct code
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
type LogFn func(o Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller)

// SetLogFn - override logging
func SetLogFn(fn LogFn) {
	if fn != nil {
		logger = fn
	}
}

// RequestConstraints - Request constraints
//type RequestConstraints interface {
//	*http.Request | Request
//}

// ResponseConstraints - Response constraints
//type ResponseConstraints interface {
//	*http.Response | *aspect.Status | int
//}

// Log - access logging.
// Header.Get(XRequestId)),
// Header.Get(XRelatesTo)),
// Header.Get(LocationHeader)
func Log(traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) {
	if logger == nil {
		return
	}
	logger(origin, traffic, start, duration, req, resp, routing, controller)
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
