package std

var (
	Origin *OriginT
)

const (
	ContentTypeText     = "text/plain charset=utf-8"
	ContentTypeBinary   = "application/octet-stream"
	ContentTypeJson     = "application/json"
	ContentTypeTextHtml = "text/html"
)

func SetOrigin(m map[string]string) *Status {
	o, status := NewOrigin(m)
	Origin = &o
	return status
}
