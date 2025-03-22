package access

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	XTo     = "X-To"
	XDomain = "X-Domain"
)

var defaultLog = func(o Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, controller Controller) {
	s := DefaultFormat(o, traffic, start, duration, req, resp, controller)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, controller Controller) string {
	newReq := BuildRequest(req)
	newResp := BuildResponse(resp)
	url, parsed := ParseURL(newReq.Host, newReq.URL)
	o.Host = Conditional(o.Host, parsed.Host)
	UpdateDefaults(&controller)

	return initFormat(o, traffic, start, duration) +
		requestFormat(o, newReq, url, parsed) +
		responseFormat(newResp) +
		controllerFormat(traffic, controller)

	/*
		s := fmt.Sprintf("{"+
			// Origin, traffic, timestamp, duration
			"\"region\":%v, "+
			"\"zone\":%v, "+
			"\"sub-zone\":%v, "+
			"\"instance-id\":%v, "+
			"\"traffic\":\"%v\", "+
			"\"start\":%v, "+
			"\"duration\":%v, "+

			// Request
			"\"request-id\":%v, "+
			"\"protocol\":%v, "+
			"\"method\":%v, "+
			"\"host\":%v, "+
			"\"uri\":%v, "+
			"\"path\":%v, "+
			"\"query\":%v, "+

			// Response
			"\"status-code\":%v, "+
			"\"encoding\":%v, "+
			"\"bytes\":%v, "+

			// Controller
			"\"timeout\":%v, "+
			"\"rate-limit\":%v, "+
			"\"rate-burst\":%v, "+
			"\"redirect\":%v, "+
			"\"cc\":%v }",

			// Origin, traffic, timestamp, duration
			JsonString(o.Region),
			JsonString(o.Zone),
			JsonString(o.SubZone),
			JsonString(o.InstanceId),
			traffic,
			FmtRFC3339Millis(start),
			strconv.Itoa(Milliseconds(duration)),

			// Request
			JsonString(newReq.Header.Get(XRequestId)),
			JsonString(newReq.Proto),
			JsonString(newReq.Method),
			JsonString(o.Host),
			JsonString(url),
			JsonString(parsed.Path),
			JsonString(parsed.Query),

			// Response
			newResp.StatusCode,
			JsonString(Encoding(newResp)),
			fmt.Sprintf("%v", newResp.ContentLength),

			// Controller
			Milliseconds(controller.Timeout),
			toInt(controller.RateLimit),
			toInt(controller.RateBurst),
			toInt(controller.Percentage),
			JsonString(controller.Code),
		)

		return s

	*/
}

func initFormat(o Origin, traffic string, start time.Time, duration time.Duration) string {
	return fmt.Sprintf("{"+
		// Origin, traffic, timestamp, duration
		"\"region\":%v, "+
		"\"zone\":%v, "+
		"\"sub-zone\":%v, "+
		"\"instance-id\":%v, "+
		"\"traffic\":\"%v\", "+
		"\"start\":%v, "+
		"\"duration\":%v, ",

		JsonString(o.Region),
		JsonString(o.Zone),
		JsonString(o.SubZone),
		JsonString(o.InstanceId),
		traffic,
		FmtRFC3339Millis(start),
		strconv.Itoa(Milliseconds(duration)))
}

func requestFormat(o Origin, newReq *http.Request, url string, parsed *Parsed) string {
	return fmt.Sprintf(
		// Request
		"\"request-id\":%v, "+
			"\"protocol\":%v, "+
			"\"method\":%v, "+
			"\"host\":%v, "+
			"\"uri\":%v, "+
			"\"path\":%v, "+
			"\"query\":%v, ",

		JsonString(newReq.Header.Get(XRequestId)),
		JsonString(newReq.Proto),
		JsonString(newReq.Method),
		JsonString(o.Host),
		JsonString(url),
		JsonString(parsed.Path),
		JsonString(parsed.Query))
}

func responseFormat(newResp *http.Response) string {
	return fmt.Sprintf(
		// Response
		"\"status-code\":%v, "+
			"\"encoding\":%v, "+
			"\"bytes\":%v, ",

		newResp.StatusCode,
		JsonString(Encoding(newResp)),
		fmt.Sprintf("%v", newResp.ContentLength),
	)
}

func controllerFormat(traffic string, controller Controller) string {
	if traffic == EgressTraffic {
		return fmt.Sprintf(
			// Controller
			"\"timeout\":%v, "+
				"\"rate-limit\":%v, "+
				"\"rate-burst\":%v, "+
				"\"redirect\":%v, "+
				"\"cc\":%v }",
			Milliseconds(controller.Timeout),
			toInt(controller.RateLimit),
			toInt(controller.RateBurst),
			toInt(controller.Percentage),
			JsonString(controller.Code),
		)
	} else {
		return fmt.Sprintf(
			// Controller
			"\"timeout\":%v, "+
				"\"rate-limit\":%v, "+
				"\"rate-burst\":%v, "+
				"\"cc\":%v }",
			Milliseconds(controller.Timeout),
			toInt(controller.RateLimit),
			toInt(controller.RateBurst),
			JsonString(controller.Code),
		)
	}
}

// Milliseconds - convert time.Duration to milliseconds
func Milliseconds(duration time.Duration) int {
	if duration <= 0 {
		return -1
	}
	return int(duration / time.Duration(1e6))
}

func BuildRequest(r any) *http.Request {
	if r == nil {
		newReq, _ := http.NewRequest("", failsafeUri, nil)
		return newReq
	}
	if req, ok := r.(*http.Request); ok {
		return req
	}
	if req, ok := r.(Request); ok {
		newReq, _ := http.NewRequest(req.Method(), req.Url(), nil)
		newReq.Header = req.Header()
		newReq.Proto = req.Protocol()
		return newReq
	}
	newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
	return newReq
}

func BuildResponse(r any) *http.Response {
	if r == nil {
		newResp := &http.Response{StatusCode: http.StatusOK}
		return newResp
	}
	if newResp, ok := r.(*http.Response); ok {
		return newResp
	}
	if sc, ok := r.(int); ok {
		return &http.Response{StatusCode: sc}
	}
	if status, ok := r.(int); ok {
		return &http.Response{StatusCode: status}
	}
	newResp := &http.Response{StatusCode: http.StatusOK}
	return newResp
}

func Encoding(resp *http.Response) string {
	encoding := ""
	if resp != nil && resp.Header != nil {
		encoding = resp.Header.Get(ContentEncoding)
	}
	// normalize encoding
	if strings.Contains(strings.ToLower(encoding), "none") {
		encoding = ""
	}
	return encoding
}

func Conditional(primary, secondary string) string {
	if len(primary) == 0 {
		return secondary
	}
	return primary
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
