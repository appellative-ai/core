package httpx

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/iox"
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

func NewResponse(statusCode int, h http.Header, content any) (resp *http.Response, status *aspect.Status) {
	resp = &http.Response{StatusCode: statusCode, ContentLength: -1, Header: h, Body: io.NopCloser(bytes.NewReader([]byte{}))}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp, aspect.NewStatus(statusCode)
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
		status = aspect.NewStatusError(aspect.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
	}
	return resp, aspect.NewStatus(statusCode)
}

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, *aspect.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, status := iox.ReadFile(uri)
	if !status.OK() {
		if strings.Contains(status.Err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, aspect.NewStatusError(aspect.StatusInvalidArgument, status.Err)
		}
		return serverErr, aspect.NewStatusError(aspect.StatusIOError, status.Err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, aspect.NewStatusError(aspect.StatusIOError, err2)
	}
	return resp1, aspect.StatusOK()

}

func NewHealthResponseOK() *http.Response {
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), healthOK)
	return resp
}

func NewNotFoundResponse() *http.Response {
	resp, _ := NewResponse(http.StatusNotFound, SetHeader(nil, ContentType, ContentTypeText), aspect.StatusNotFound().String())
	return resp
}
