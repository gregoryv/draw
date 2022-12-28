package design

import (
	"io"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/shape"
)

// NewDiagram returns a diagram with present font and padding values.
//
// TODO: size and padding affects eg. records, but is related to the
// styling
func NewDiagram() *Diagram {
	return &Diagram{
		Style: draw.NewStyle(),
	}
}

// Diagram is a generic SVG image with box related styling
type Diagram struct {
	draw.SVG
	shape.Aligner
	draw.Style

	Caption *shape.Label
}

// Note is a convenience method to add a shape.Note to a diagram
func (d *Diagram) Note(v string) *shape.Adjuster {
	return d.Place(shape.NewNote(v))
}

// Place adds the shape to the diagram returning an adjuster for
// positioning.
func (d *Diagram) Place(s ...shape.Shape) *shape.Adjuster {
	for _, s := range s {
		d.applyStyle(s)
		d.Append(s)
	}
	adj := shape.NewAdjuster(s...)
	adj.Spacing = d.Style.Spacing
	return adj
}

// PlaceGrid place all the shapes into a grid starting at X,Y
// position. Row height is adapted to highest element.
func (d *Diagram) PlaceGrid(cols, X, Y int, s ...shape.Shape) {
	row := make([]shape.Shape, cols)
	var x, y int
	var h shape.Shape
	for i, s := range s {
		switch {
		case i == 0:
			d.Place(s).At(X, Y)
		case y == 0:
			d.Place(s).RightOf(row[x-1])
		default:
			d.Place(s).Below(h)
			d.VAlignCenter(row[x], s)
		}
		row[x] = s
		x++
		if x == cols {
			x = 0
			y++
			h = highest(row...)
		}
	}
}

func highest(s ...shape.Shape) shape.Shape {
	var h int
	var r shape.Shape
	for _, s := range s {
		if s.Height() > h {
			h = s.Height()
			r = s
		}
	}
	return r
}

// LinkAll places arrows between each shape, s0->s1->...->sn
func (d *Diagram) LinkAll(s ...shape.Shape) {
	for i, next := range s[1:] {
		d.Place(shape.NewArrowBetween(s[i], next))
	}
}

// Link places an arrow with a optional label above it between the two
// shapes.
func (d *Diagram) Link(from, to shape.Shape, txt ...string) (lnk *shape.Line, label *shape.Label) {
	lnk = shape.NewArrowBetween(from, to)
	d.Place(lnk)

	if len(txt) > 0 {
		label = shape.NewLabel(txt[0])

		x, y := lnk.CenterPosition()
		d.Place(label).At(x, y)

		d.HAlignCenter(lnk, label)

		// find center of arrow
		dir := lnk.Direction()
		if dir == shape.DirectionLeft || dir == shape.DirectionRight {
			d.VAlignCenter(lnk, label)
			shape.Move(label, 0, label.Font.LineHeight/2)
		}
		if dir == shape.DirectionUp || dir == shape.DirectionDown {
			shape.Move(label, label.Pad.Left, 0)
		}
	}
	return
}

func (d *Diagram) applyStyle(s interface{}) {
	if s, ok := s.(shape.HasFont); ok {
		s.SetFont(d.Font)
	}
	if s, ok := s.(shape.HasTextPad); ok {
		s.SetTextPad(d.TextPad)
	}
}

// SaveAs saves the diagram to filename as SVG
func (d *Diagram) SaveAs(filename string) error {
	return saveAs(d, d.Style, filename)
}

// Inline returns rendered SVG with inlined style
func (d *Diagram) Inline() string {
	return inline(d, d.Style)
}

// String returns rendered SVG
func (d *Diagram) String() string { return toString(d) }

func (d *Diagram) WriteSVG(w io.Writer) error {
	if d.Width() == 0 && d.Height() == 0 {
		d.AdaptSize()
	}
	if d.Caption != nil {
		margin := 30
		x := (d.Width() - d.Caption.Width()) / 2
		if x < 0 {
			x = 0
		}
		d.Place(d.Caption).At(x, d.Height()+margin)
		d.AdaptSize()
		d.SetHeight(d.Height() + d.Caption.Font.Height/2) // Fit protruding letters like 'g'
	}
	return d.SVG.WriteSVG(w)
}

// AdaptSize adapts the diagram size to the shapes inside it so all
// are visible. Returns the new width and height
func (d *Diagram) AdaptSize() (int, int) {
	for _, s := range d.Content {
		s, ok := s.(shape.Shape)
		if !ok {
			continue
		}
		x, y := s.Position()
		if s, ok := s.(*shape.Line); ok {
			x = min(s.Start.X, s.End.X)
			y = min(s.Start.Y, s.End.Y)
		}
		w := x + s.Width()
		if w > d.Width() {
			d.SetWidth(w)
		}
		h := y + s.Height()
		if h > d.Height() {
			d.SetHeight(h)
		}
	}
	d.SetWidth(d.Width() + 1)   // Fixes right most pixels not visible
	d.SetHeight(d.Height() + 1) // Fixes bottom pixels not visible
	return d.Width(), d.Height()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SetCaption adds a caption to the bottom of the diagram.
func (d *Diagram) SetCaption(txt string) {
	l := shape.NewLabel(txt)
	l.SetClass("caption")
	d.Caption = l
}
