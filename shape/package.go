package shape

import (
	"io"
)

type SvgWriter interface {
	WriteSvg(io.Writer) error
}

type Font struct {
	Height     int
	Width      int
	LineHeight int
}

type Padding struct {
	Left, Top, Right, Bottom int
}

func boxHeight(font Font, pad Padding, lines int) int {
	return (pad.Top + pad.Bottom) + lines*font.LineHeight
}

func boxWidth(font Font, pad Padding, txt string) int {
	return pad.Left + font.Width*len(txt) + pad.Right
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
