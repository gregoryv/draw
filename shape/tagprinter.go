package shape

import (
	"fmt"
	"io"
)

func newTagPrinter(w io.Writer) (*tagPrinter, *error) {
	tag := &tagPrinter{w: w}
	return tag, &tag.err
}

type printFunc func(...interface{})
type printFfunc func(string, ...interface{})

// tagPrinter is used to write out svg tags and if an error has
// occured previously it's methods do nothing (the nexus).
type tagPrinter struct {
	w   io.Writer
	err error
}

func (ec *tagPrinter) printf(format string, args ...interface{}) {
	if ec.err != nil {
		return
	}
	_, ec.err = fmt.Fprintf(ec.w, format, args...)
}

func (ec *tagPrinter) print(args ...interface{}) {
	if ec.err != nil {
		return
	}
	_, ec.err = fmt.Fprint(ec.w, args...)
}

func (ec *tagPrinter) Write(b []byte) (int, error) {
	if ec.err != nil {
		return 0, ec.err
	}
	return ec.w.Write(b)
}
