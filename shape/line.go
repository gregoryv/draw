package shape

import (
	"github.com/gregoryv/draw/xy"
)

func NewLine(x1, y1 int, x2, y2 int) *Line {
	return &Line{
		Arrow{
			Start: xy.Point{x1, y1},
			End:   xy.Point{x2, y2},
			class: "line",
		},
	}
}

type Line struct {
	Arrow
}
