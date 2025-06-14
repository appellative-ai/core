package messaging

var (
	Origin *OriginT
)

func SetOrigin(m map[string]string) *Status {
	o, status := NewOrigin(m)
	Origin = &o
	return status
}
