package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewDot() *Dot {
	return &Dot{
		Radius: 6,
		class:  "dot",
	}
}

type Dot struct {
	x, y   int
	Radius int
	class  string
}

func (d *Dot) String() string {
	return fmt.Sprintf("Dot")
}

func (d *Dot) Position() (int, int) {
	return d.x, d.y
}

func (d *Dot) SetX(x int) { d.x = x }
func (d *Dot) SetY(y int) { d.y = y }
func (d *Dot) Width() int {
	return d.Radius * 2
}
func (d *Dot) Height() int           { return d.Width() }
func (d *Dot) Direction() Direction  { return DirectionRight }
func (d *Dot) SetClass(class string) { d.class = class }

func (d *Dot) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := d.Position()
	x += d.Radius
	y += d.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		d.class, x, y, d.Radius,
	)
	return *err
}

func (d *Dot) Edge(start xy.Point) xy.Point {
	return boxEdge(start, d)
}
