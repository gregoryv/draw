package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewDot() *Dot {
	return &Dot{
		Radius: 6,
		class:  "dot",
	}
}

type Dot struct {
	pos    xy.Position
	Radius int
	class  string
}

func (c *Dot) String() string {
	return fmt.Sprintf("Dot")
}

func (c *Dot) Position() (int, int) {
	return c.pos.XY()
}

func (c *Dot) SetX(x int) { c.pos.X = x }
func (c *Dot) SetY(y int) { c.pos.Y = y }
func (c *Dot) Width() int {
	return c.Radius * 2
}
func (c *Dot) Height() int           { return c.Width() }
func (c *Dot) Direction() Direction  { return LR }
func (c *Dot) SetClass(class string) { c.class = class }

func (c *Dot) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagPrinter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius,
	)
	return *err
}

func (c *Dot) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
