package httpx

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

const (
	internalError  = "Internal Error"
	gatewayTimeout = "Gateway Timeout"
	fileScheme     = "file"
)

var (
	Client          = http.DefaultClient
	serverResponse  = serverErrorResponse()
	timeoutResponse = gatewayTimeoutResponse()
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
	// panic if req or URL is nil - should be resolved during testing
	if req.URL.Scheme == fileScheme {
		return NewResponseFromUri(req.URL)
	}
	resp, err = Client.Do(req)
	// catch *url.Error - can be a connectivity or a context deadline exceeded error
	if err != nil {
		if urlErr, ok := any(err).(*url.Error); ok {
			if urlErr.Timeout() {
				return timeoutResponse, err
			}
		}
		resp = serverResponse
	}
	return
}

// DoWithTimeout - process an HTTP request with a timeout and optional Exchange
func DoWithTimeout(req *http.Request, timeout time.Duration, ex Exchange) (resp *http.Response, err error) {
	if ex == nil {
		ex = Do
	}
	if timeout <= 0 {
		return ex(req)
	}
	ctx, cancel := NewContext(timeout)
	defer cancel()
	resp, err = ex(req.Clone(ctx))
	if err == nil {
		err = TransformBody(resp)
	}
	return
}

func serverErrorResponse() *http.Response {
	resp := new(http.Response)
	resp.StatusCode = http.StatusInternalServerError
	resp.Status = internalError
	resp.Body = EmptyReader
	return resp
}

func gatewayTimeoutResponse() *http.Response {
	resp := new(http.Response)
	resp.StatusCode = http.StatusGatewayTimeout
	resp.Status = gatewayTimeout
	resp.Body = EmptyReader
	return resp
}

/*
	if req == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, errors.New("invalid argument : request is nil")
	}
*/

// convert deadline exceeded error into a Gateway Timeout
/*
	if req.Context() != nil && req.Context().Err() == context.DeadlineExceeded {
		resp = gatewayTimeoutResponse()
		err = nil
	} else {
		if resp == nil {
			resp = serverErrorResponse(err)
		}
	}

*/
