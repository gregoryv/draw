package shape

import (
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewState(title string) *State {
	return &State{
		Title: title,
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultTextPad,
		class: "state",
	}
}

type State struct {
	x, y  int
	Title string

	Font  draw.Font
	Pad   draw.Padding
	class string
}

func (r *State) Position() (x int, y int) { return r.x, r.y }
func (r *State) SetX(x int)               { r.x = x }
func (r *State) SetY(y int)               { r.y = y }
func (r *State) Direction() Direction     { return DirectionRight }
func (r *State) SetClass(c string)        { r.class = c }

func (r *State) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.x, r.y, r.Width(), r.Height())
	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *State) title() *Label {
	return &Label{
		x:     r.x + r.Pad.Left,
		y:     r.y + r.Pad.Top/2,
		Font:  r.Font,
		text:  r.Title,
		class: "state-title",
	}
}

func (r *State) SetFont(f draw.Font)         { r.Font = f }
func (r *State) SetTextPad(pad draw.Padding) { r.Pad = pad }

func (r *State) Height() int {
	return boxHeight(r.Font, r.Pad, 1)
}

func (r *State) Width() int {
	return boxWidth(r.Font, r.Pad, r.Title)
}

// Edge returns intersecting position of a line starting at start and
// pointing to the rect center.
func (r *State) Edge(start xy.Point) xy.Point {
	return boxEdge(start, r)
}
