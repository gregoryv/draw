// Package shape provides various SVG shapes
package shape

import (
	"math"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func boxHeight(font draw.Font, pad draw.Padding, lines int) int {
	return (pad.Top + pad.Bottom) + lines*font.LineHeight
}

func boxWidth(font draw.Font, pad draw.Padding, txt string) int {
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
	SetFont(draw.Font)
}

type HasTextPad interface {
	SetTextPad(draw.Padding)
}
