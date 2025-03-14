package http

import (
	"net/http"
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
