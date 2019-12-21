package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewExitDot() *ExitDot {
	return &ExitDot{
		Radius: 10,
		class:  "exit",
	}
}

type ExitDot struct {
	x, y   int
	Radius int
	class  string
}

func (c *ExitDot) String() string {
	return fmt.Sprintf("ExitDot")
}

func (c *ExitDot) Position() (int, int) { return c.x, c.y }

func (c *ExitDot) SetX(x int) { c.x = x }
func (c *ExitDot) SetY(y int) { c.y = y }
func (c *ExitDot) Width() int {
	// If the style shanges the width will be slightly off, no biggy.
	return c.Radius*2 + 4
}
func (c *ExitDot) Height() int           { return c.Width() }
func (c *ExitDot) Direction() Direction  { return LR }
func (c *ExitDot) SetClass(class string) { c.class = class }

func (c *ExitDot) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	x, y := c.Position()
	x += c.Radius
	y += c.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x+2, y+2, c.Radius,
	)
	w.Printf(
		`<circle class="%s-dot" cx="%v" cy="%v" r="%v" />\n`,
		c.class, x+2, y+2, c.Radius-4,
	)

	return *err
}

func (c *ExitDot) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
