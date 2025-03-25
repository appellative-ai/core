package httpx

import "io"

var (
	emptyReader = new(nilReader)
)

type nilReader struct{}

func (r *nilReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (r *nilReader) Close() error {
	return nil
}
