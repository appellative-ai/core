package access

import (
	"fmt"
	"github.com/behavioral-ai/core/fmtx"
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

var defaultLog = func(o *Origin, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) {
	s := DefaultFormat(o, traffic, start, duration, route, req, resp, thresholds)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o *Origin, traffic string, start time.Time, duration time.Duration, route string, req any, resp any, thresholds Threshold) string {
	newReq := BuildRequest(req)
	newResp := BuildResponse(resp)
	url, parsed := ParseURL(newReq.Host, newReq.URL)
	//o.Host = Conditional(o.Host, parsed.Host)
	//UpdateDefaults(&controller)

	return initFormat(o, traffic, start, duration, route) +
		requestFormat(newReq, url, parsed) +
		responseFormat(newResp) +
		thresholdsFormat(traffic, thresholds)

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

func initFormat(o *Origin, traffic string, start time.Time, duration time.Duration, route string) string {
	if o == nil {
		return fmt.Sprintf("{"+
			// Traffic, timestamp, duration
			"\"traffic\":\"%v\", "+
			"\"start\":%v, "+
			"\"duration\":%v, "+
			"\"route\":%v, ",

			traffic,
			fmtx.FmtRFC3339Millis(start),
			strconv.Itoa(fmtx.Milliseconds(duration)),
			fmtx.JsonString(route))
	} else {
		return fmt.Sprintf("{"+
			// Origin, traffic, timestamp, duration
			"\"region\":%v, "+
			"\"zone\":%v, "+
			"\"sub-zone\":%v, "+
			"\"instance-id\":%v, "+
			"\"traffic\":\"%v\", "+
			"\"start\":%v, "+
			"\"duration\":%v, "+
			"\"route\":%v, ",

			fmtx.JsonString(o.Region),
			fmtx.JsonString(o.Zone),
			fmtx.JsonString(o.SubZone),
			fmtx.JsonString(o.InstanceId),
			//fmtx.JsonString(o.Route),
			traffic,
			fmtx.FmtRFC3339Millis(start),
			strconv.Itoa(fmtx.Milliseconds(duration)),
			fmtx.JsonString(route))
	}
}

func requestFormat(newReq *http.Request, url string, parsed *Parsed) string {
	return fmt.Sprintf(
		// Request
		"\"request-id\":%v, "+
			"\"protocol\":%v, "+
			"\"method\":%v, "+
			"\"host\":%v, "+
			"\"uri\":%v, "+
			"\"path\":%v, "+
			"\"query\":%v, ",

		fmtx.JsonString(newReq.Header.Get(XRequestId)),
		fmtx.JsonString(newReq.Proto),
		fmtx.JsonString(newReq.Method),
		fmtx.JsonString(newReq.Host),
		fmtx.JsonString(url),
		fmtx.JsonString(parsed.Path),
		fmtx.JsonString(parsed.Query))
}

func responseFormat(newResp *http.Response) string {
	return fmt.Sprintf(
		// Response
		"\"status-code\":%v, "+
			"\"encoding\":%v, "+
			"\"bytes\":%v, ",

		newResp.StatusCode,
		fmtx.JsonString(Encoding(newResp)),
		fmt.Sprintf("%v", newResp.ContentLength),
	)
}

func thresholdsFormat(traffic string, thresholds Threshold) string {
	if traffic == EgressTraffic {
		return fmt.Sprintf(
			"\"timeout\":%v, "+
				"\"rate-limit\":%v, "+
				"\"redirect\":%v } ",
			fmtx.Milliseconds(thresholds.timeout()),
			thresholds.rateLimit(),
			thresholds.redirect(),
		)
	} else {
		return fmt.Sprintf(
			"\"timeout\":%v, "+
				"\"rate-limit\":%v, "+
				"\"redirect\":%v } ",
			fmtx.Milliseconds(thresholds.timeout()),
			thresholds.rateLimit(),
			thresholds.redirect(),
		)
	}
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
