package httpx

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/iox"
	"io"
)

const (
	eofError = "EOF"
)

func Content[T any](body io.Reader) (t T, status *aspect.Status) {
	if body == nil {
		return t, aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	status = aspect.StatusOK()
	switch p := any(&t).(type) {
	case *[]byte:
		var buf []byte
		buf, status = iox.ReadAll(body, nil)
		if !status.OK() {
			return
		}
		if len(buf) == 0 {
			return t, aspect.StatusNotFound()
		}
		*p = buf
	case *string:
		var buf []byte
		buf, status = iox.ReadAll(body, nil)
		if !status.OK() {
			return
		}
		if len(buf) == 0 {
			return t, aspect.StatusNotFound()
		}
		*p = string(buf)
	default:
		err := json.NewDecoder(body).Decode(p)
		if err != nil {
			// If the error is "EOF", then the body was empty. If the error is "unexpected EOF", then the body has content
			// but the EOF was reached when more JSON content was expected.
			if err.Error() == eofError {
				status = aspect.StatusNoContent()
			} else {
				status = aspect.NewStatusError(aspect.StatusJsonDecodeError, err)
			}
		}
	}
	return
}
