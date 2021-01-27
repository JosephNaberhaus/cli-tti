package util

import (
	"bytes"
	"io"
)

// A writer that will accumulate in an internal buffer of all the bytes it writes
type AccumulatorWriter struct {
	written bytes.Buffer
	Writer io.Writer
}

func (a *AccumulatorWriter) Write(p []byte) (n int, err error) {
	n, err = a.Writer.Write(p)
	if err != nil {
		return n, err
	}

	a.written.Write(p)

	return n, nil
}

func (a *AccumulatorWriter) Written() []byte {
	return a.written.Bytes()
}
