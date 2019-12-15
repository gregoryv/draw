package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
)

func NewTriangle(x, y int, class string) *Triangle {
	return &Triangle{
		pos:   xy.Position{x, y},
		class: class,
	}
}

type Triangle struct {
	pos   xy.Position
	class string
}

func (tri *Triangle) String() string {
	return fmt.Sprintf("triangle at %v", tri.pos)
}

func (t *Triangle) Position() (int, int) { return t.pos.XY() }
func (t *Triangle) SetX(x int)           { t.pos.X = x }
func (t *Triangle) SetY(y int)           { t.pos.Y = y }
func (t *Triangle) Width() int           { return 8 }
func (t *Triangle) Height() int          { return 4 }
func (t *Triangle) Direction() Direction { return LR }
func (t *Triangle) SetClass(c string)    { t.class = c }

func (tri *Triangle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	// the path is drawn as if it points straight to the right
	w.printf(`<path class="%s" d="M%v,%v l-8,-4 l 0,8 Z" />`,
		tri.class, tri.pos.X, tri.pos.Y)
	return *err
}
