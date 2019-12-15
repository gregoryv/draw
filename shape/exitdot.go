package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
)

func NewExitDot() *ExitDot {
	return &ExitDot{
		Radius: 10,
		class:  "exit",
	}
}

type ExitDot struct {
	pos    xy.Position
	Radius int
	class  string
}

func (c *ExitDot) String() string {
	return fmt.Sprintf("ExitDot")
}

func (c *ExitDot) Position() (int, int) {
	return c.pos.XY()
}

func (c *ExitDot) SetX(x int) { c.pos.X = x }
func (c *ExitDot) SetY(y int) { c.pos.Y = y }
func (c *ExitDot) Width() int {
	// If the style shanges the width will be slightly off, no biggy.
	stroke := 2
	return (c.Radius+stroke)*2 - 2
}
func (c *ExitDot) Height() int           { return c.Width() }
func (c *ExitDot) Direction() Direction  { return LR }
func (c *ExitDot) SetClass(class string) { c.class = class }

func (c *ExitDot) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius,
	)
	w.printf(
		`<circle class="%s-dot" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x, y, c.Radius-4,
	)

	return *err
}

func (c *ExitDot) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
