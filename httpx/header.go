package httpx

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	ContentTypeJson     = "application/json"
	ContentType         = "Content-Type"
	ContentEncoding     = "Content-Encoding"
	AcceptEncoding      = "Accept-Encoding"
	AcceptEncodingValue = "gzip, deflate, br"
	ContentLength       = "Content-Length"
	ContentEncodingGzip = "gzip"
	ContentTypeTextHtml = "text/html"
	ContentTypeText     = "text/plain charset=utf-8"
	ContentLocation     = "Content-Location"
	ExchangeOverride    = "X-Exchange-Override"
	ContentResolver     = "X-Content-Resolver"
	ResolverSeparator   = "->"
)

func SetHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Set(name, value)
	return h
}

func SetHeaders(w http.ResponseWriter, headers any) {
	if headers == nil {
		return
	}
	if pairs, ok := headers.([]Attr); ok {
		for _, pair := range pairs {
			w.Header().Set(strings.ToLower(pair.Key), pair.Value)
		}
		return
	}
	if h, ok := headers.(http.Header); ok {
		for k, v := range h {
			if len(v) > 0 {
				w.Header().Set(strings.ToLower(k), v[0])
			}
		}
	}
}

func Copy(h http.Header) http.Header {
	h2 := make(http.Header)
	if h != nil {
		for k, v := range h {
			fmt.Printf("header: %v %v\n", k, v)
			h2.Set(k, v[0])
		}
	}
	return h2
}
