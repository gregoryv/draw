package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/nexus"
)

func NewTriangle() *Triangle {
	return &Triangle{
		class:  "triangle",
		width:  8,
		height: 8,
	}
}

type Triangle struct {
	x, y          int
	width, height int
	class         string
}

func (t *Triangle) String() string {
	return fmt.Sprintf("triangle at %v,%v", t.x, t.y)
}

// fixme should point to TopLeft corner
func (t *Triangle) Position() (int, int) { return t.x, t.y }
func (t *Triangle) SetX(x int)           { t.x = x }
func (t *Triangle) SetY(y int)           { t.y = y }
func (t *Triangle) Width() int           { return t.width }
func (t *Triangle) Height() int          { return t.height }
func (t *Triangle) Direction() Direction { return DirectionRight }
func (t *Triangle) SetClass(c string)    { t.class = c }

func (t *Triangle) WriteSVG(out io.Writer) error {
	p, err := nexus.NewPrinter(out)

	w, h := t.width, t.height
	p.Printf(`<path class="%s" d="M%v,%v l%v,%v l %v,%v Z" />`,
		t.class, t.x, t.y, -w, -h/2, 0, h)
	return *err
}
