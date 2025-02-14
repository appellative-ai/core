package jsonx

import (
	"bytes"
	"encoding/json"
	"io"
)

// NewReadCloser - create an io.ReadCloser from a type
func NewReadCloser(v any) (io.ReadCloser, int64, error) {
	if v == nil {
		return io.NopCloser(bytes.NewReader([]byte{})), 0, nil
	}
	buf, status := json.Marshal(v)
	if status != nil {
		return nil, 0, status
	}
	return io.NopCloser(bytes.NewReader(buf)), int64(len(buf)), nil
}
