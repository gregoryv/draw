package draw

import (
	"fmt"
	"io"
)

// NewTagWriter returns a writer nexus. The referenced error is set by
// all methods if an error occurs.
func NewTagWriter(w io.Writer) (*TagWriter, *error) {
	t := &TagWriter{w: w}
	return t, &t.err
}

type TagWriter struct {
	w   io.Writer
	err error
}

// Print prints arguments using the underlying writer. Does nothing if
// TagWriter has failed.
func (t *TagWriter) Print(args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprint(t.w, args...)
}

// Printf prints a formated string using the underlying writer. Does
// nothing if TagWriter has failed.
func (t *TagWriter) Printf(format string, args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprintf(t.w, format, args...)
}

// Write writes the bytes using the underlying writer. Does nothing if
// TagWriter has failed.
func (t *TagWriter) Write(b []byte) (int, error) {
	if t.err != nil {
		return 0, t.err
	}
	return t.w.Write(b)
}
