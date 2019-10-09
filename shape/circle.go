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
	topLeft xy.Position
	center  xy.Position
	Radius  int
	class   string
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle")
}

func (c *Circle) Position() (int, int) { return c.topLeft.XY() }
func (c *Circle) SetX(x int) {
	c.topLeft.X = x
	c.center.X = x + c.Radius
}
func (c *Circle) SetY(y int) {
	c.topLeft.Y = y
	c.center.Y = y + c.Radius
}
func (c *Circle) Width() int            { return c.Radius * 2 }
func (c *Circle) Height() int           { return c.Width() }
func (c *Circle) Direction() Direction  { return LR }
func (c *Circle) SetClass(class string) { c.class = class }

func (c *Circle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, c.center.X, c.center.Y, c.Radius,
	)
	return *err
}
