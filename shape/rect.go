package shape

import (
	"io"
	"math"
	"strings"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewRect(title string) *Rect {
	return &Rect{
		Title: title,
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultTextPad,
		class: "rect",
	}
}

type Rect struct {
	x, y  int
	Title string

	Font  draw.Font
	Pad   draw.Padding
	class string

	width, height int

	textAlign string
}

func (r *Rect) Position() (x int, y int) { return r.x, r.y }
func (r *Rect) SetX(x int)               { r.x = x }
func (r *Rect) SetY(y int)               { r.y = y }
func (r *Rect) Direction() Direction     { return DirectionRight }
func (r *Rect) SetClass(c string)        { r.class = c }

func (r *Rect) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	w.Printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.x, r.y, r.Width(), r.Height())
	w.Printf("\n")
	r.title().WriteSVG(w)
	return *err
}

func (r *Rect) title() *Label {
	return &Label{
		x:     r.x + r.Pad.Left,
		y:     r.y + r.Pad.Top/2,
		Font:  r.Font,
		text:  r.Title,
		class: r.class + "-title",
	}
}

func (r *Rect) SetFont(f draw.Font)         { r.Font = f }
func (r *Rect) SetTextPad(pad draw.Padding) { r.Pad = pad }

func (r *Rect) Height() int {
	if r.height > 0 {
		return r.height
	}
	lines := strings.Count(r.Title, "\n") + 1
	return boxHeight(r.Font, r.Pad, lines)
}

func (r *Rect) Width() int {
	if r.width > 0 {
		return r.width
	}
	return boxWidth(r.Font, r.Pad, longestLine(r.Title))
}

func (r *Rect) SetWidth(w int)  { r.width = w }
func (r *Rect) SetHeight(h int) { r.height = h }

// Edge returns intersecting position of a line starting at start and
// pointing to the rect center.
func (r *Rect) Edge(start xy.Point) xy.Point {
	return boxEdge(start, r)
}

type Box interface {
	// Position returns the xy position of the top left corner.
	Position() (x int, y int)
	Width() int
	Height() int
}

func boxEdge(start xy.Point, r Box) xy.Point {
	x, y := r.Position()
	center := xy.Point{
		x + r.Width()/2,
		y + r.Height()/2,
	}
	l1 := xy.Line{start, center}

	var (
		d      float64 = math.MaxFloat64
		pos    xy.Point
		lowY   = y + r.Height()
		rightX = x + r.Width()
		top    = xy.NewLine(x, y, rightX, y)
		left   = xy.NewLine(x, y, x, lowY)
		right  = xy.NewLine(rightX, y, rightX, lowY)
		bottom = xy.NewLine(x, lowY, rightX, lowY)
	)

	for _, side := range []*xy.Line{top, left, right, bottom} {
		p, err := l1.IntersectSegment(side)
		if err != nil {
			continue
		}
		dist := start.Distance(p)
		if dist < d {
			pos = p
			d = dist
		}
	}
	return pos
}
