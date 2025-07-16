package jsonx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/appellative-ai/core/iox"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// NewConstraints - constraints
//type NewConstraints interface {
//	string | *url.URL | []byte | iox.Reader | iox.ReadCloser
//}

const (
	eofError = "EOF"
)

func decodeStatus(err error) error {
	if err == nil || err.Error() == "" {
		return nil
	}
	// If the error is "EOF", then the body was empty. If the error is "unexpected EOF", then the body has content
	// but the EOF was reached when more JSON content was expected.
	if err.Error() == eofError {
		return errors.New("error: no content") //aspect.StatusNoContent()
	}
	return err
}

//type NewConstraints interface {
//	string | []byte | *url.URL | *httpx.Request | *httpx.Response | interface{ iox.Reader } | interface{ iox.ReadCloser }
//}

// New - create a new type from JSON content, supporting: string, *url.URL, []byte, io.Reader, io.ReadCloser
func New[T any](v any, h http.Header) (t T, status error) {
	var buf []byte

	if v == nil {
		return t, errors.New("error: value parameter is nil")
	}
	switch ptr := v.(type) {
	case string:
		//if isStatusURL(ptr) {
		//	return t, NewStatusFrom(ptr)
		//}
		buf, status = iox.ReadFileWithEncoding(ptr, h)
		if status != nil {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case *url.URL:
		//if isStatusURL(ptr.String()) {
		//	return t, NewStatusFrom(ptr.String())
		//}
		buf, status = iox.ReadFileWithEncoding(ptr.String(), h)
		if status != nil {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case []byte:
		buf, status = iox.Decode(ptr, h)
		if status != nil {
			return
		}
		err := json.Unmarshal(buf, &t)
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case io.Reader:
		reader, status0 := iox.NewEncodingReader(ptr, h)
		if status0 != nil {
			return t, status0 //status0.AddLocation()
		}
		err := json.NewDecoder(reader).Decode(&t)
		_ = reader.Close()
		//if err != nil {
		//	return t, decodeStatus(err)
		//	}
		return t, decodeStatus(err)
	case io.ReadCloser:
		reader, status0 := iox.NewEncodingReader(ptr, h)
		if status0 != nil {
			return t, status0 //status0.AddLocation()
		}
		err := json.NewDecoder(reader).Decode(&t)
		_ = reader.Close()
		_ = ptr.Close()
		//if err != nil {
		//	return t, decodeStatus(err)
		//}
		return t, decodeStatus(err)
	case *http.Request:
		return New[T](ptr.Body, h)
	case *http.Response:
		return New[T](ptr.Body, h)
	default:
		return t, errors.New(fmt.Sprintf("error: invalid type [%v]", reflect.TypeOf(v)))
	}
}

/*
	case *httpx.Response:
		if ptr1, ok := any(&t).(*[]byte); ok {
			buf, status = ReadAll(ptr.Body,h)
			if !status.OK() {
				return
			}
			*ptr1 = buf
			return t, StatusOK()
		}
		err := jsonx.NewDecoder(ptr.Body).Decode(&t)
		_ = ptr.Body.Close()
		if err != nil {
			return t, NewStatus(StatusJsonDecodeError, err)
		}
		return t, StatusOK()
	case *httpx.Request:
		if ptr1, ok := any(&t).(*[]byte); ok {
			buf, status = ReadAll(ptr.Body)
			if !status.OK() {
				return
			}
			*ptr1 = buf
			return t, StatusOK()
		}
		err := jsonx.NewDecoder(ptr.Body).Decode(&t)
		_ = ptr.Body.Close()
		if err != nil {
			return t, NewStatus(StatusJsonDecodeError, err)
		}
		return t, StatusOK()

*/
