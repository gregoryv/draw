package shape

import (
	"io"
)

type SvgWriter interface {
	WriteSvg(io.Writer) error
}

type Font struct {
	Height int
	Width  int
}

type Padding struct {
	Left, Top, Right, Bottom int
}

type ErrCollector struct {
	err error
}

func (ec *ErrCollector) Last(any interface{}, err error) {
	ec.Err(err)
}

func (ec *ErrCollector) Err(err error) {
	if ec.err != nil {
		return
	}
	ec.err = err
}

func (ec *ErrCollector) First() error {
	return ec.err
}
