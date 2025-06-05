package messaging

var (
	Origin OriginT
)

func SetOrigin(m map[string]string, collective, domain string) (status *Status) {
	Origin, status = NewOrigin(m, collective, domain)
	return status
}
