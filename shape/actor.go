package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
)

// NewActor returns a new actor with a default height.
func NewActor() *Actor {
	return &Actor{
		height: 35,
		class:  "actor",
	}
}

type Actor struct {
	pos    xy.Position // top left
	height int
	class  string
}

func (c *Actor) String() string        { return fmt.Sprintf("Actor") }
func (c *Actor) Position() (int, int)  { return c.pos.XY() }
func (c *Actor) SetX(x int)            { c.pos.X = x }
func (c *Actor) SetY(y int)            { c.pos.Y = y }
func (c *Actor) Width() int            { return c.rad() * 4 }
func (c *Actor) Height() int           { return c.height }
func (c *Actor) SetHeight(h int)       { c.height = h }
func (c *Actor) rad() int              { return c.height / 6 }
func (c *Actor) Direction() Direction  { return LR }
func (c *Actor) SetClass(class string) { c.class = class }

func (c *Actor) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	x, y := c.Position()
	r := c.rad()
	d := r * 2
	// head
	w.printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />`,
		c.class, x+d, y+r, r,
	)
	w.print("\n")
	// body
	w.printf(`<path class="%s" d="M%v,%v l 0,%v m -%v,-%v l %v,0 m -%v,%v l -%v,%v m %v,-%v l %v,%v Z" />`,
		c.class, x+d, y+d, r*3, d, d, c.Width(), d, d, d, d, d, d, d, d)
	w.print("\n")
	return *err
}

func (c *Actor) Edge(start xy.Position) xy.Position {
	return boxEdge(start, c)
}
