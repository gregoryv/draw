package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
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
	pos    xy.Position
	width  int
	height int
	class  string
}

func (d *Diamond) String() string {
	return fmt.Sprintf("Diamond at %v", d.pos)
}

func (d *Diamond) Position() (int, int) { return d.pos.XY() }
func (d *Diamond) SetX(x int)           { d.pos.X = x }
func (d *Diamond) SetY(y int)           { d.pos.Y = y - d.height/2 }
func (d *Diamond) Width() int           { return d.width }
func (d *Diamond) Height() int          { return d.height }
func (d *Diamond) Direction() Direction { return LR }
func (d *Diamond) SetClass(c string)    { d.class = c }

func (d *Diamond) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagPrinter(out)
	x, y := d.pos.XY()
	y += d.height / 2
	w2 := d.width / 2
	h2 := d.height / 2
	// the path is drawn from left to right
	w.Printf(`<path class="%s" d="M%v,%v l %v,%v %v,%v %v,%v %v,%v" />`,
		d.class, x, y, w2, -h2, w2, h2, -w2, h2, -w2, -h2)
	return *err
}

func (d *Diamond) Edge(start xy.Position) xy.Position {
	return boxEdge(start, d)
}
