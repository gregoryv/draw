package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
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

func (e *ExitDot) String() string {
	return fmt.Sprintf("ExitDot")
}

func (e *ExitDot) Position() (x int, y int) { return e.x, e.y }
func (e *ExitDot) SetX(x int) { e.x = x }
func (e *ExitDot) SetY(y int) { e.y = y }

func (e *ExitDot) Width() int {
	// If the style shanges the width will be slightly off, no biggy.
	return e.Radius*2 + 4
}
func (e *ExitDot) Height() int           { return e.Width() }
func (e *ExitDot) Direction() Direction  { return DirectionRight }
func (e *ExitDot) SetClass(class string) { e.class = class }

func (e *ExitDot) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := e.Position()
	x += e.Radius
	y += e.Radius
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />\n`,
		e.class, x+2, y+2, e.Radius,
	)
	w.Printf(
		`<circle class="%s-dot" cx="%v" cy="%v" r="%v" />\n`,
		e.class, x+2, y+2, e.Radius-4,
	)

	return *err
}

func (e *ExitDot) Edge(start xy.Point) xy.Point {
	return boxEdge(start, e)
}
