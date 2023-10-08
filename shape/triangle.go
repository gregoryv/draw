package shape

import (
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

func (t *Triangle) Position() (x int, y int) { return t.x, t.y }
func (t *Triangle) SetX(x int)               { t.x = x }
func (t *Triangle) SetY(y int)               { t.y = y }
func (t *Triangle) Width() int               { return t.width }
func (t *Triangle) Height() int              { return t.height }
func (t *Triangle) Direction() Direction     { return DirectionRight }
func (t *Triangle) SetClass(c string)        { t.class = c }

func (t *Triangle) WriteSVG(out io.Writer) error {
	p, err := nexus.NewPrinter(out)

	w, h := t.width, t.height
	w2 := w / 2
	/*
	      +
	     / \
	    /   \
	   +-----+
	*/
	p.Printf(`<path class="%s" d="M%v,%v l%v,%v l %v,%v Z" />`,
		t.class, t.x, t.y+h, w2, -h, w2, h)
	return *err
}
