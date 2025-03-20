package http

import (
	"context"
	"crypto/tls"
	"errors"
	http2 "net/http"
	"time"
)

const (
	internalError           = "Internal Error"
	fileScheme              = "file"
	contextDeadlineExceeded = "context deadline exceeded"
)

type Exchange func(r *http2.Request) (*http2.Response, error)

var (
	Client = http2.DefaultClient
)

func init() {
	t, ok := http2.DefaultTransport.(*http2.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		Client = &http2.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		Client = &http2.Client{Transport: http2.DefaultTransport, Timeout: time.Second * 5}
	}
}

// Do - process an HTTP request, checking for file:// scheme
func Do(req *http2.Request) (resp *http2.Response, err error) {
	if req == nil {
		return &http2.Response{StatusCode: http2.StatusInternalServerError}, errors.New("invalid argument : request is nil")
	}
	if req.URL.Scheme == fileScheme {
		return NewResponseFromUri(req.URL)
	}
	resp, err = Client.Do(req)
	if err != nil {
		//if urlErr, ok := any(err).(*url.Error); ok {
		//}
		// catch connectivity error, even with a valid URL
		if resp == nil {
			resp = serverErrorResponse()
		}
		// check for an error of deadline exceeded
		if req.Context() != nil && req.Context().Err() == context.DeadlineExceeded {
			resp.StatusCode = http2.StatusGatewayTimeout
			err = errors.New(contextDeadlineExceeded)
		}
	}
	return
}

func serverErrorResponse() *http2.Response {
	resp := new(http2.Response)
	resp.StatusCode = http2.StatusInternalServerError
	resp.Status = internalError
	return resp
}
