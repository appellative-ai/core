package iox

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ReadAll - read the body with a Status
func ReadAll(body io.Reader, h http.Header) ([]byte, *aspect.Status) {
	if body == nil {
		return nil, aspect.StatusOK()
	}
	if rc, ok := any(body).(io.ReadCloser); ok {
		defer func() {
			err := rc.Close()
			if err != nil {
				fmt.Printf("error: iox.ReadCloser.Close() [%v]", err)
			}
		}()
	}
	reader, status := NewEncodingReader(body, h)
	if !status.OK() {
		return nil, status.AddLocation()
	}
	buf, err := io.ReadAll(reader)
	_ = reader.Close()
	if err != nil {
		return nil, aspect.NewStatusError(aspect.StatusIOError, err)
	}
	return buf, aspect.StatusOK()
}

func ValidateUri(uri string) *aspect.Status {
	if len(uri) == 0 {
		return aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URI is empty"))
	}
	if !strings.HasPrefix(uri, fileScheme) {
		return aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New(fmt.Sprintf("error: URI is not of scheme file: %v", uri)))
	}
	////if !isJsonURL(uri) {
	//	return aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URI is not a JSON file"))
	//}
	return aspect.StatusOK()
}
