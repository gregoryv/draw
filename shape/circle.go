package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewCircle(radius int) *Circle {
	return &Circle{
		Radius: radius,
		class:  "circle",
	}
}

type Circle struct {
	xy.Point
	Radius int
	class  string
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle")
}

func (c *Circle) Width() int {
	stroke := 1
	return (c.Radius+stroke)*2 - 2
}
func (c *Circle) Height() int           { return c.Width() }
func (c *Circle) Direction() Direction  { return DirectionRight }
func (c *Circle) SetClass(class string) { c.class = class }

func (c *Circle) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius,
	)
	return *err
}

func (c *Circle) Edge(start xy.Point) xy.Point {
	return boxEdge(start, c)
}
