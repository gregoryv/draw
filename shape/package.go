// Package shape provides various SVG shapes
package shape

import (
	"math"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

type Padding = draw.Padding
type Font = draw.Font

var (
	DefaultFont    = draw.DefaultFont
	DefaultPad     = draw.DefaultPad
	DefaultTextPad = draw.DefaultTextPad
	DefaultSpacing = draw.DefaultSpacing
)

func boxHeight(font Font, pad Padding, lines int) int {
	return (pad.Top + pad.Bottom) + lines*font.LineHeight
}

func boxWidth(font Font, pad Padding, txt string) int {
	return pad.Left + font.TextWidth(txt) + pad.Right
}

func intAbs(v int) int {
	return int(math.Abs(float64(v)))
}

type Edge interface {
	// Edge returns the intersecting position to a shape from start
	// position.
	Edge(start xy.Point) xy.Point
}

type HasFont interface {
	SetFont(Font)
}

type HasTextPad interface {
	SetTextPad(Padding)
}
