package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

// NewActor returns a new actor with a default height.
func NewActor() *Actor {
	return &Actor{
		height: 35,
		class:  "actor",
	}
}

type Actor struct {
	x, y   int // top left
	height int
	class  string
}

func (a *Actor) String() string        { return fmt.Sprintf("Actor") }
func (a *Actor) Position() (int, int)  { return a.x, a.y }
func (a *Actor) SetX(x int)            { a.x = x }
func (a *Actor) SetY(y int)            { a.y = y }
func (a *Actor) Width() int            { return a.rad() * 4 }
func (a *Actor) Height() int           { return a.height }
func (a *Actor) SetHeight(h int)       { a.height = h }
func (a *Actor) rad() int              { return a.height / 6 }
func (a *Actor) Direction() Direction  { return DirectionRight }
func (a *Actor) SetClass(class string) { a.class = class }

func (a *Actor) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x, y := a.Position()
	r := a.rad()
	d := r * 2
	// head
	w.Printf(
		`<circle class="%s" cx="%v" cy="%v" r="%v" />`,
		a.class, x+d, y+r, r,
	)
	w.Print("\n")
	// body
	w.Printf(`<path class="%s" d="M%v,%v l 0,%v m -%v,-%v l %v,0 m -%v,%v l -%v,%v m %v,-%v l %v,%v Z" />`,
		a.class, x+d, y+d, r*3, d, d, a.Width(), d, d, d, d, d, d, d, d)
	w.Print("\n")
	return *err
}

func (a *Actor) Edge(start xy.Position) xy.Position {
	return boxEdge(start, a)
}
