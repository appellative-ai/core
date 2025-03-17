package http

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

const (
	internalError           = "Internal Error"
	fileScheme              = "file"
	contextDeadlineExceeded = "context deadline exceeded"
)

type Exchange func(r *http.Request) (*http.Response, error)

var (
	Client = http.DefaultClient
)

func init() {
	t, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		Client = &http.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		Client = &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 5}
	}
}

// Do - process an HTTP request, checking for file:// scheme
func Do(req *http.Request) (resp *http.Response, err error) {
	if req == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, errors.New("invalid argument : request is nil")
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
			resp.StatusCode = http.StatusGatewayTimeout
			err = errors.New(contextDeadlineExceeded)
		}
	}
	return
}

func serverErrorResponse() *http.Response {
	resp := new(http.Response)
	resp.StatusCode = http.StatusInternalServerError
	resp.Status = internalError
	return resp
}
