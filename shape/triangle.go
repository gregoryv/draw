package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewTriangle(x, y int, class string) *Triangle {
	return &Triangle{
		Pos:   xy.Position{x, y},
		Class: class,
	}
}

type Triangle struct {
	Pos   xy.Position
	Class string
}

func (tri *Triangle) String() string {
	return fmt.Sprintf("triangle at %v", tri.Pos)
}

func (tri *Triangle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	// the path is drawn as if it points straight to the right
	w.printf(`<path class="%s" d="M%v,%v l-8,-4 l 0,8 Z" />`,
		tri.Class, tri.Pos.X, tri.Pos.Y)
	return *err
}

func (tri *Triangle) Direction() Direction {
	return LR
}

func (tri *Triangle) Height() int {
	return 4
}

func (tri *Triangle) Width() int {
	return 8
}

func (tri *Triangle) Position() (int, int) {
	return tri.Pos.XY()
}

func (tri *Triangle) SetX(x int) {
	tri.Pos.X = x
}

func (tri *Triangle) SetY(y int) {
	tri.Pos.Y = y
}
