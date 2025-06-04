package messaging

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Content -
type Content struct {
	Fragment string // returned on a Get
	Type     string // Content-Type
	Value    any
}

func (c Content) String() string {
	return fmt.Sprintf("fragment: %v type: %v value: %v", c.Fragment, c.Type, c.Value != nil)
}

func (c Content) Valid(contentType string) bool {
	return c.Value != nil && c.Type == contentType
}

// Unmarshal - []byte -> string, []byte, io.Reader, type via json.Unmarshal
func Unmarshal[T any](ct *Content) (t T, status *Status) {
	var body []byte
	var ok bool

	if ct == nil {
		return t, NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: content is nil")))
	}
	if ct.Type == "" || ct.Value == nil {
		return t, NewStatus(http.StatusNoContent, errors.New(fmt.Sprintf("error: content type is empty, or content value is nil")))
	}
	if body, ok = ct.Value.([]byte); !ok {
		return t, NewStatus(StatusInvalidContent, fmt.Sprintf("error: content type is not []byte"))
	}
	if len(body) == 0 {
		return t, StatusOK()
	}
	switch ptr := any(&t).(type) {
	case *string:
		if ct.Type != ContentTypeText && ct.Type != ContentTypeTextHtml {
			return t, NewStatus(StatusInvalidContent, fmt.Sprintf("error: content type %v invalid for string", ct.Type))
		}
		*ptr = string(body)
	case *[]byte:
		if ct.Type != ContentTypeBinary {
			return t, NewStatus(StatusInvalidContent, fmt.Sprintf("error: content type %v invalid for []byte", ct.Type))
		}
		*ptr = body
	default:
		if ct.Type != ContentTypeJson {
			return t, NewStatus(StatusInvalidContent, fmt.Sprintf("error: content type %v invalid for %v", ct.Type, reflect.TypeOf(t)))
		}
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, NewStatus(StatusJsonDecodeError, errors.New(fmt.Sprintf("error: JSON unmarshalling %v", err)))
		}
	}
	return t, StatusOK()
}

// Marshal -  type -> []byte | io.Reader
func Marshal[T any](ct *Content) (t T, status *Status) {
	var buf []byte

	if ct == nil {
		return t, NewStatus(http.StatusBadRequest, errors.New(fmt.Sprintf("error: content is nil")))
	}
	if ct.Type == "" || ct.Value == nil {
		return t, NewStatus(http.StatusNoContent, errors.New(fmt.Sprintf("error: content type is empty, or content value is nil")))
	}
	switch ptr := ct.Value.(type) {
	case string:
		buf = []byte(ptr)
	case []byte:
		buf = ptr
	default:
		var err error
		buf, err = json.Marshal(ptr)
		if err != nil {
			return t, NewStatus(StatusJsonEncodeError, err)
		}
	}
	if len(buf) == 0 {
		return t, NewStatus(http.StatusNoContent, errors.New("content value is empty"))
	}
	switch ptr := any(&t).(type) {
	case *[]byte:
		*ptr = buf
		return t, StatusOK()
	case *io.Reader:
		*ptr = bytes.NewReader(buf)
		return t, StatusOK()
	default:
	}
	return t, NewStatus(StatusInvalidContent, fmt.Sprintf("error: content type %v invalid for %v", ct.Type, reflect.TypeOf(t)))
}
