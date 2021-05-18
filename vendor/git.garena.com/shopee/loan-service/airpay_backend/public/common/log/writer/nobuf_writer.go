package writer

import "io"

func NewNoBufWriter(w io.Writer, size int) BufferedWriter {
	return &noBufWriter{w}
}

type noBufWriter struct {
	wr io.Writer
}

func (w *noBufWriter) Write(p []byte) (n int, err error) {
	return w.wr.Write(p)
}

func (w *noBufWriter) Flush() error {
	return nil
}
