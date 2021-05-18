package writer

import "time"

const (
	defaultBufSize     = 4096
	defaultFlushPeriod = 1 * time.Second
)

type BufferedWriter interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
