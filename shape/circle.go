package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewCircle(radius int) *Circle {
	return &Circle{
		Radius: radius,
		class:  "circle",
	}
}

type Circle struct {
	pos    xy.Position // top left
	Radius int
	class  string
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle")
}

func (c *Circle) Position() (int, int) {
	return c.pos.XY()
}

func (c *Circle) SetX(x int)            { c.pos.X = x }
func (c *Circle) SetY(y int)            { c.pos.Y = y }
func (c *Circle) Width() int            { return c.Radius * 2 }
func (c *Circle) Height() int           { return c.Width() }
func (c *Circle) Direction() Direction  { return LR }
func (c *Circle) SetClass(class string) { c.class = class }

func (c *Circle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius,
	)
	return *err
}

func (c *Circle) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
