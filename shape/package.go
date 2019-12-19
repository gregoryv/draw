// Package shape provides various SVG shapes
package shape

import (
	"math"

	"github.com/gregoryv/draw/xy"
)

type Padding struct {
	Left, Top, Right, Bottom int
}

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
	Edge(start xy.Position) xy.Position
}

type HasFont interface {
	SetFont(Font)
}

type HasTextPad interface {
	SetTextPad(Padding)
}
