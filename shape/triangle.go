package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
)

func NewTriangle() *Triangle {
	return &Triangle{
		class: "triangle",
	}
}

type Triangle struct {
	x, y  int
	class string
}

func (t *Triangle) String() string {
	return fmt.Sprintf("triangle at %v,%v", t.x, t.y)
}

func (t *Triangle) Position() (int, int) { return t.x, t.y }
func (t *Triangle) SetX(x int)           { t.x = x }
func (t *Triangle) SetY(y int)           { t.y = y }
func (t *Triangle) Width() int           { return 8 }
func (t *Triangle) Height() int          { return 4 }
func (t *Triangle) Direction() Direction { return RightDir }
func (t *Triangle) SetClass(c string)    { t.class = c }

func (t *Triangle) WriteSVG(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	// the path is drawn as if it points straight to the right
	w.Printf(`<path class="%s" d="M%v,%v l-8,-4 l 0,8 Z" />`,
		t.class, t.x, t.y)
	return *err
}
