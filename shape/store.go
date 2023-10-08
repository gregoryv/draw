package shape

import (
	"fmt"
	"io"
	"strings"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewStore(title string) *Store {
	return &Store{
		Title: title,
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultTextPad,
		class: "store",
	}
}

type Store struct {
	x, y  int
	Title string

	Font  draw.Font
	Pad   draw.Padding
	class string

	width, height int
}

func (r *Store) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *Store) Position() (x int, y int) { return r.x, r.y }
func (r *Store) SetX(x int)               { r.x = x }
func (r *Store) SetY(y int)               { r.y = y }
func (r *Store) Direction() Direction     { return DirectionRight }
func (r *Store) SetClass(c string)        { r.class = c }

func (r *Store) WriteSVG(out io.Writer) error {
	y := r.y
	x := r.x
	w, err := nexus.NewPrinter(out)
	top := NewLine(x, y, x+r.Width(), y)
	top.SetClass(r.class)
	top.WriteSVG(w)

	y = y + r.Height()
	bottom := NewLine(x, y, x+r.Width(), y)
	bottom.SetClass(r.class)
	bottom.WriteSVG(w)

	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *Store) title() *Label {
	return &Label{
		x:     r.x + r.Pad.Left,
		y:     r.y + r.Pad.Top/2,
		Font:  r.Font,
		text:  r.Title,
		class: r.class + "-title",
	}
}

func (r *Store) SetFont(f draw.Font)         { r.Font = f }
func (r *Store) SetTextPad(pad draw.Padding) { r.Pad = pad }

func (r *Store) Height() int {
	if r.height > 0 {
		return r.height
	}
	return boxHeight(r.Font, r.Pad, strings.Count(r.Title, "\n")+1)
}

func (r *Store) Width() int {
	if r.width > 0 {
		return r.width
	}
	return boxWidth(r.Font, r.Pad, r.Title)
}

func (r *Store) SetWidth(w int)  { r.width = w }
func (r *Store) SetHeight(h int) { r.height = h }

func (r *Store) Edge(start xy.Point) xy.Point {
	return boxEdge(start, r)
}
