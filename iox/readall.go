package iox

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"io"
	"net/http"
	"strings"
)

// ReadAll - read the body with a Status
func ReadAll(body io.Reader, h http.Header) ([]byte, *core.Status) {
	if body == nil {
		return nil, core.StatusOK()
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
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	return buf, core.StatusOK()
}

func ValidateUri(uri string) *core.Status {
	if len(uri) == 0 {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URI is empty"))
	}
	if !strings.HasPrefix(uri, fileScheme) {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: URI is not of scheme file: %v", uri)))
	}
	////if !isJsonURL(uri) {
	//	return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URI is not a JSON file"))
	//}
	return core.StatusOK()
}
