package util

import "io"

type LineWriter struct {
	lines int
	inner io.Writer
}

func NewLineWriter(w io.Writer) *LineWriter {
	return &LineWriter{
		inner: w,
	}
}

func (l *LineWriter) Reset() {
	l.lines = 0
}

func (l *LineWriter) Lines() int {
	return l.lines
}

func (l *LineWriter) Write(p []byte) (n int, err error) {
	_, err = l.inner.Write(p)
	if err != nil {
		return 0, err
	}
	for _, b := range p {
		if b == 0x0a {
			l.lines++
			continue
		}
	}
	return len(p), nil
}
