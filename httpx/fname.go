package httpx

import (
	"github.com/appellative-ai/core/iox"
	"io"
	"net/http"
)

func fileName(uri any) string {
	return iox.FileName(uri)
}

func readFile(uri any) ([]byte, error) {
	return iox.ReadFile(uri)
}

func readAll(body io.Reader) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return io.ReadAll(body)
}

func newEncodingWriter(w io.Writer, h http.Header) (iox.EncodingWriter, error) {
	return iox.NewEncodingWriter(w, h)
}
