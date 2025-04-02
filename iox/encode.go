package iox

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
)

func EncodeContent(r *http.Request, content []byte) (buff []byte, encoding string, err error) {
	if r == nil || r.Header == nil {
		return nil, "", errors.New("request or request header is nil")
	}
	enc := acceptEncoding(r.Header)
	if !strings.Contains(enc, GzipEncoding) {
		return nil, "", nil
	}
	buf := new(bytes.Buffer)
	w, err1 := NewEncodingWriter(buf, r.Header)
	if err1 != nil {
		return nil, "", err
	}
	cnt, err2 := w.Write(content)
	err3 := w.Close()
	if err2 != nil || cnt != len(content) {
		return nil, "", err2
	}
	if err3 != nil {
		return nil, "", err3
	}
	return buf.Bytes(), GzipEncoding, nil
}
