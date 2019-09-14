// Package shape provides various SVG shapes
package shape

import (
	"io"
	"math"
)

type SvgWriter interface {
	WriteSvg(io.Writer) error
}

type SvgWriterShape interface {
	SvgWriter
	Shape
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

func intAbs(v int) int {
	return int(math.Abs(float64(v)))
}

type HasFont interface {
	SetFont(Font)
}

type HasTextPad interface {
	SetTextPad(Padding)
}
