package test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"os"
	"testing"
)

/*
const (
	fileExistsError = "The system cannot find the file specified"
	fileScheme      = "file"
)

// NewResponse - read an HTTP response given a URL
func NewResponse(uri any) (*http.Response, *aspect.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: "Internal Error"}

	if uri == nil {
		return serverErr, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, status := io.ReadFile(uri)
	if !status.OK() {
		if strings.Contains(status.Err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found"}, aspect.NewStatusError(aspect.StatusInvalidArgument, status.Err)
		}
		return serverErr, aspect.NewStatusError(aspect.StatusIOError, status.Err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, aspect.NewStatusError(aspect.StatusIOError, err2)
	}
	return resp1, aspect.StatusOK()
}

*/

// NewResponseTest - used from test packages
func NewResponseTest(uri any, t *testing.T) *http.Response {
	resp, status := httpx.NewResponseFromUri(uri)
	if status.OK() {
		return resp
	}
	t.Errorf("ReadResponse() err = %v", status.Err.Error())
	return &http.Response{StatusCode: http.StatusTeapot}
}

func writeValues(buf *bytes.Buffer, name string, values []string) {
	for _, value := range values {
		buf.WriteString(fmt.Sprintf("%v: %v\n", name, value))
	}
}

func writeHeader(buf *bytes.Buffer, resp *http.Response) {
	for name, values := range resp.Header {
		writeValues(buf, name, values)
	}
}

func WriteResponse(url string, resp *http.Response) *aspect.Status {
	if url == "" || resp == nil {
		return aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: url is empty or response is nil"))
	}
	var buf bytes.Buffer

	// Write status line
	buf.WriteString(fmt.Sprintf("%v %v\n", resp.Proto, resp.Status))

	// Write header
	writeHeader(&buf, resp)
	buf.WriteString("\n")

	// Write content
	buf1, status := iox.ReadAll(resp.Body, nil)
	if !status.OK() {
		return status
	}
	count, err := buf.Write(buf1)
	if err != nil {
		return aspect.NewStatusError(aspect.StatusIOError, err)
	}
	if count != len(buf1) {
		return aspect.NewStatusError(aspect.StatusIOError, errors.New("error: writing bytes"))
	}

	// Create filename and write file
	fname := iox.FileName(url)
	// 0666 is read only
	err = os.WriteFile(fname, buf.Bytes(), 0777)
	if err != nil {
		return aspect.NewStatusError(aspect.StatusIOError, err)
	}
	return aspect.StatusOK()
}
