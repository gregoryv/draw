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

func (c *Circle) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, c.center.X, c.center.Y, c.Radius,
	)
	return *err
}
