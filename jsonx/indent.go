package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/appellative-ai/core/iox"
	"io"
	"net/http"
)

func Indent(body io.ReadCloser, h http.Header, prefix, indent string) (io.ReadCloser, error) {
	var buf bytes.Buffer

	if body == nil {
		return nil, errors.New("error: body is nil")
	}
	buf2, status := iox.ReadAll(body, h)
	if status != nil {
		return nil, status
	}
	err := json.Indent(&buf, buf2, prefix, indent)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
