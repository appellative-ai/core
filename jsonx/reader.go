package jsonx

import (
	"bytes"
	"io"
)

// NewReadCloser - create an io.ReadCloser from a type
func NewReadCloser(v any) (io.ReadCloser, int64, *aspect.Status) {
	if v == nil {
		return io.NopCloser(bytes.NewReader([]byte{})), 0, aspect.StatusOK()
	}
	buf, status := Marshal(v)
	if !status.OK() {
		return nil, 0, status
	}
	return io.NopCloser(bytes.NewReader(buf)), int64(len(buf)), aspect.StatusOK()
}
