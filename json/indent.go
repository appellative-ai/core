package json

import (
	"bytes"
	"encoding/json"
	"errors"
	iox "github.com/behavioral-ai/core/io"
	"io"
	"net/http"
)

func Indent(body io.ReadCloser, h http.Header, prefix, indent string) (io.ReadCloser, error) {
	var buf bytes.Buffer

	if body == nil {
		return nil, errors.New("error: body is nil") //aspect.NewStatusError(aspect.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	buf2, status := iox.ReadAll(body, h)
	if status != nil {
		return nil, status
	}
	err := json.Indent(&buf, buf2, prefix, indent)
	if err != nil {
		return nil, err //aspect.NewStatusError(aspect.StatusJsonDecodeError, err)
	}
	return io.NopCloser(bytes.NewReader(buf.Bytes())), nil //aspect.StatusOK()
}
