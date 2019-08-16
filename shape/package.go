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
