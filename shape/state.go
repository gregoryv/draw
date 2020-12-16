package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewState(title string) *State {
	return &State{
		Title: title,
		Font:  DefaultFont,
		Pad:   DefaultTextPad,
		class: "state",
	}
}

type State struct {
	X, Y  int
	Title string

	Font  Font
	Pad   Padding
	class string
}

func (r *State) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *State) Position() (int, int) { return r.X, r.Y }
func (r *State) SetX(x int)           { r.X = x }
func (r *State) SetY(y int)           { r.Y = y }
func (r *State) Direction() Direction { return DirectionRight }
func (r *State) SetClass(c string)    { r.class = c }

func (r *State) WriteSVG(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X, r.Y, r.Width(), r.Height())
	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *State) title() *Label {
	return &Label{
		x:     r.X + r.Pad.Left,
		y:     r.Y + r.Pad.Top/2,
		Font:  r.Font,
		Text:  r.Title,
		class: "state-title",
	}
}

func (r *State) SetFont(f Font)         { r.Font = f }
func (r *State) SetTextPad(pad Padding) { r.Pad = pad }

func (r *State) Height() int {
	return boxHeight(r.Font, r.Pad, 1)
}

func (r *State) Width() int {
	return boxWidth(r.Font, r.Pad, r.Title)
}

// Edge returns intersecting position of a line starting at start and
// pointing to the rect center.
func (r *State) Edge(start xy.Position) xy.Position {
	return boxEdge(start, r)
}
