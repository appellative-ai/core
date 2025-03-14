package http

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	io2 "github.com/behavioral-ai/core/io"
	"github.com/behavioral-ai/core/messaging"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const (
	fileExistsError = "The system cannot find the file specified"
)

var (
	healthOK = []byte("{\n \"status\": \"up\"\n}")
	//healthLength = int64(len(healthOK))
)

func NewResponse(statusCode int, h http.Header, content any) (resp *http.Response, status *messaging.Status) {
	resp = &http.Response{StatusCode: statusCode, ContentLength: -1, Header: h, Body: io.NopCloser(bytes.NewReader([]byte{}))}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp, messaging.NewStatus(statusCode)
	}
	switch ptr := (content).(type) {
	case []byte:
		if len(ptr) > 0 {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader(ptr))
		}
	case string:
		if ptr != "" {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
		}
	case error:
		if ptr.Error() != "" {
			resp.ContentLength = int64(len(ptr.Error()))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
		}
	default:
		status = messaging.NewStatusError(messaging.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))), "")
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
	}
	return resp, messaging.NewStatus(statusCode)
}

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, *messaging.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, messaging.NewStatusError(messaging.StatusInvalidArgument, errors.New("error: URL is nil"), "")
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, messaging.NewStatusError(messaging.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, err := io2.ReadFile(uri)
	if err != nil {
		if strings.Contains(err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, messaging.NewStatusError(messaging.StatusInvalidArgument, err, "")
		}
		return serverErr, messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, messaging.NewStatusError(messaging.StatusIOError, err2, "")
	}
	return resp1, messaging.StatusOK()

}

func NewHealthResponseOK() *http.Response {
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), healthOK)
	return resp
}

func NewNotFoundResponse() *http.Response {
	resp, _ := NewResponse(http.StatusNotFound, SetHeader(nil, ContentType, ContentTypeText), messaging.StatusNotFound().String())
	return resp
}
