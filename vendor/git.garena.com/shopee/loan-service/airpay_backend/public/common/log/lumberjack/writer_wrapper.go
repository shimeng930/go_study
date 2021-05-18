package lumberjack

import (
	"io"

	"git.garena.com/shopee/loan-service/airpay_backend/public/common/log/writer"
)

type WriterWrapper func(w io.Writer) writer.BufferedWriter

var (
	defaultWriterWrapper func(w io.Writer) writer.BufferedWriter
)

func init() {
	WithDoubleBufWrapper(4 * 1024)
}

func WithNoBufWrapper() {
	defaultWriterWrapper = func(w io.Writer) writer.BufferedWriter { return writer.NewNoBufWriter(w, 0) }
}

func WithBufioWrapper(size int) {
	defaultWriterWrapper = func(w io.Writer) writer.BufferedWriter { return writer.NewBufioWriterSize(w, size) }
}

func WithDoubleBufWrapper(size int) {
	defaultWriterWrapper = func(w io.Writer) writer.BufferedWriter { return writer.NewDoubleBufWriterSize(w, size) }
}
