package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewCircle(radius int) *Circle {
	return &Circle{
		Radius: radius,
		class:  "circle",
	}
}

type Circle struct {
	x, y   int // top left
	Radius int
	class  string
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle")
}

func (c *Circle) Position() (int, int) { return c.x, c.y }

func (c *Circle) SetX(x int) { c.x = x }
func (c *Circle) SetY(y int) { c.y = y }
func (c *Circle) Width() int {
	stroke := 1
	return (c.Radius+stroke)*2 - 2
}
func (c *Circle) Height() int           { return c.Width() }
func (c *Circle) Direction() Direction  { return DirectionRight }
func (c *Circle) SetClass(class string) { c.class = class }

func (c *Circle) WriteSVG(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius,
	)
	return *err
}

func (c *Circle) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
