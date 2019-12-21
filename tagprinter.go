package draw

import (
	"fmt"
	"io"
)

// NewTagPrinter returns a printer nexus.
func NewTagWriter(w io.Writer) (*TagWriter, *error) {
	tag := &TagWriter{w: w}
	return tag, &tag.err
}

type printFunc func(...interface{})
type printFfunc func(string, ...interface{})

// TagPrinter is used to write out svg tags and if an error has
// occured previously it's methods do nothing.
type TagWriter struct {
	w   io.Writer
	err error
}

func (ec *TagWriter) Printf(format string, args ...interface{}) {
	if ec.err != nil {
		return
	}
	_, ec.err = fmt.Fprintf(ec.w, format, args...)
}

func (ec *TagWriter) Print(args ...interface{}) {
	if ec.err != nil {
		return
	}
	_, ec.err = fmt.Fprint(ec.w, args...)
}

func (ec *TagWriter) Write(b []byte) (int, error) {
	if ec.err != nil {
		return 0, ec.err
	}
	return ec.w.Write(b)
}
