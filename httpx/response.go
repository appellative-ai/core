package httpx

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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

// TransformBody - read the body and create a new []byte buffer reader
func TransformBody(resp *http.Response) error {
	if resp == nil || resp.Body == nil {
		return nil
	}
	buf, err := io.ReadAll(resp.Body)
	if err == nil {
		resp.ContentLength = int64(len(buf))
		resp.Body = io.NopCloser(bytes.NewReader(buf))
	}
	return err
}

func NewResponse(statusCode int, h http.Header, content any) (resp *http.Response) {
	resp = &http.Response{StatusCode: statusCode, ContentLength: -1, Header: h, Body: EmptyReader}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp
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
		err := errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr)))
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(err.Error())))}
	}
	return resp
}

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, error) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, errors.New("error: URL is nil")
	}
	buf, err := iox.ReadFile(uri)
	if err != nil {
		if strings.Contains(err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, err
		}
		return serverErr, err
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, err2
	}
	return resp1, nil

}

func NewHealthResponseOK() *http.Response {
	return NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), healthOK)
}

func NewNotFoundResponse() *http.Response {
	return NewResponse(http.StatusNotFound, SetHeader(nil, ContentType, ContentTypeText), "Not Found")
}
