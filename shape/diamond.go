package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewDiamond() *Diamond {
	return &Diamond{
		width:  12,
		height: 8,
		class:  "diamond",
	}
}

func NewDecision() *Diamond {
	return &Diamond{
		width:  20,
		height: 20,
		class:  "decision",
	}
}

type Diamond struct {
	x, y   int
	width  int
	height int
	class  string
}

func (d *Diamond) String() string {
	return fmt.Sprintf("Diamond at %v,%v", d.x, d.y)
}

func (d *Diamond) Position() (int, int) { return d.x, d.y }
func (d *Diamond) SetX(x int)           { d.x = x }
func (d *Diamond) SetY(y int)           { d.y = y }
func (d *Diamond) Width() int           { return d.width }
func (d *Diamond) Height() int          { return d.height }
func (d *Diamond) Direction() Direction { return DirectionRight }
func (d *Diamond) SetClass(c string)    { d.class = c }

func (d *Diamond) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := d.Position()
	w2 := d.width / 2
	h2 := d.height / 2
	// the path is drawn from left to right
	w.Printf(`<path class="%s" d="M%v,%v l %v,%v %v,%v %v,%v %v,%v" />`,
		d.class, x, y+h2, w2, -h2, w2, h2, -w2, h2, -w2, -h2)
	return *err
}

func (d *Diamond) Edge(start xy.Point) xy.Point {
	return boxEdge(start, d)
}
