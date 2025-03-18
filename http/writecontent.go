package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	iox "github.com/behavioral-ai/core/io"
	"io"
	"reflect"
	"strings"
)

const (
	jsonToken = "json"
)

func writeContent(w io.Writer, content any, contentType string) (length int64, status *aspect.Status) {
	var err error
	var cnt int

	if content == nil {
		return 0, aspect.StatusOK()
	}
	switch ptr := (content).(type) {
	case []byte:
		cnt, err = w.Write(ptr)
	case string:
		cnt, err = w.Write([]byte(ptr))
	case error:
		cnt, err = w.Write([]byte(ptr.Error()))
	case io.Reader:
		var buf []byte
		var err1 error

		buf, err1 = iox.ReadAll(ptr, nil)
		if err1 != nil {
			status = aspect.NewStatusError(aspect.StatusIOError, err)
			return 0, status.AddLocation()
		}
		status = aspect.StatusOK()
		cnt, err = w.Write(buf)
	case io.ReadCloser:
		var buf []byte
		var err1 error

		buf, err1 = iox.ReadAll(ptr, nil)
		_ = ptr.Close()
		if err1 != nil {
			status = aspect.NewStatusError(aspect.StatusIOError, err)
			return 0, status.AddLocation()
		}
		status = aspect.StatusOK()
		cnt, err = w.Write(buf)
	default:
		if strings.Contains(contentType, jsonToken) {
			var buf []byte

			buf, err = json.Marshal(content)
			if err != nil {
				status = aspect.NewStatusError(aspect.StatusJsonEncodeError, err)
				if !status.OK() {
					return
				}
			}
			cnt, err = w.Write(buf)
		} else {
			return 0, aspect.NewStatusError(aspect.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		}
	}
	if err != nil {
		return 0, aspect.NewStatusError(aspect.StatusIOError, err)
	}
	return int64(cnt), aspect.StatusOK()
}
