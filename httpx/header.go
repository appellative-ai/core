package httpx

import (
	"github.com/behavioral-ai/core/aspect"
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

func forwardDefaults(dest http.Header, src http.Header) http.Header {
	if dest == nil {
		dest = make(http.Header)
	}
	if src == nil {
		return dest
	}
	// TODO : add other default headers
	dest.Set(aspect.XRequestId, src.Get(aspect.XRequestId))
	dest.Set(aspect.XRelatesTo, src.Get(aspect.XRelatesTo))
	dest.Set(aspect.XFrom, src.Get(aspect.XFrom))
	// Verify
	//dest.Set(aspect.XTest, src.Get(aspect.XTest))
	return dest
}

// Forward - forward headers
func Forward(dest http.Header, src http.Header, names ...string) http.Header {
	dest = forwardDefaults(dest, src)
	if src == nil {
		return dest
	}
	for _, name := range names {
		dest.Set(name, src.Get(name))
	}
	return dest
}

// GetContentType - return the content type header value
func GetContentType(headers any) string {
	if pairs, ok := headers.([]Attr); ok {
		for _, pair := range pairs {
			if pair.Key == ContentType {
				return pair.Value
			}
		}
		return ""
	}
	if h, ok := headers.(http.Header); ok {
		return h.Get(ContentType)
	}
	return ""
}

// SetHeaders - set the headers into an HTTP response writer
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

func AddHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Add(name, value)
	return h
}

func SetHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Set(name, value)
	return h
}
