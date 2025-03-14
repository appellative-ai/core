package io

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ReadAll - read the body with a Status
func ReadAll(body io.Reader, h http.Header) ([]byte, error) {
	if body == nil {
		return nil, nil //aspect.StatusOK()
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
	if status != nil {
		return nil, status //status.AddLocation()
	}
	buf, err := io.ReadAll(reader)
	_ = reader.Close()
	if err != nil {
		return nil, err //aspect.NewStatusError(aspect.StatusIOError, err)
	}
	return buf, nil
}

func ValidateUri(uri string) error {
	if len(uri) == 0 {
		return errors.New("error: URI is empty") //aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URI is empty"))
	}
	if !strings.HasPrefix(uri, fileScheme) {
		return errors.New(fmt.Sprintf("error: URI is not of scheme file: %v", uri)) //aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New(fmt.Sprintf("error: URI is not of scheme file: %v", uri)))
	}
	////if !isJsonURL(uri) {
	//	return aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: URI is not a JSON file"))
	//}
	return nil
}
