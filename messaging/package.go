package messaging

var (
	Origin OriginT
)

func SetOrigin(m map[string]string) (status *Status) {
	Origin, status = NewOrigin(m)
	return status
}
