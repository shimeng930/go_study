package writer

import (
	"bufio"
	"io"
	"sync"
	"time"
)

func NewBufioWriterSize(w io.Writer, size int) BufferedWriter {
	if size <= 0 {
		size = defaultBufSize
	}
	bw := &bufioWriter{
		wr: bufio.NewWriterSize(w, size),
	}
	go bw.flushPeriodically()
	return bw
}

type bufioWriter struct {
	mu sync.Mutex
	wr *bufio.Writer
}

func (w *bufioWriter) Flush() error {
	return w.flushInner()
}

func (w *bufioWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.wr.Write(p)
}

func (w *bufioWriter) flushPeriodically() {
	ticker := time.NewTicker(defaultFlushPeriod)
	defer ticker.Stop()
	for {
		<-ticker.C
		w.flushInner()
	}
}

func (w *bufioWriter) flushInner() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.wr.Flush()
}
