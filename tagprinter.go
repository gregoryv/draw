package draw

import (
	"fmt"
	"io"
)

// NewTagPrinter returns a printer nexus.
func NewTagWriter(w io.Writer) (*TagWriter, *error) {
	t := &TagWriter{w: w}
	return t, &t.err
}

// TagPrinter is used to write out svg tags and if an error has
// occured previously it's methods do nothing.
type TagWriter struct {
	w   io.Writer
	err error
}

func (t *TagWriter) Printf(format string, args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprintf(t.w, format, args...)
}

func (t *TagWriter) Print(args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprint(t.w, args...)
}

func (t *TagWriter) Write(b []byte) (int, error) {
	if t.err != nil {
		return 0, t.err
	}
	return t.w.Write(b)
}
