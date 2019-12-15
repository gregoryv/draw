package design

import (
	"io"

	"github.com/gregoryv/go-design/shape"
)

// NewDiagram returns a diagram with present font and padding values.
//
// TODO: size and padding affects eg. records, but is related to the
// styling
func NewDiagram() Diagram {
	return Diagram{
		Style: shape.NewStyle(nil),
	}
}

// Diagram is a generic SVG image with box related styling
type Diagram struct {
	shape.Svg
	shape.Aligner
	shape.Style

	Caption *shape.Label
}

// Place adds the shape to the diagram returning an adjuster for
// positioning.
func (diagram *Diagram) Place(s ...shape.Shape) *shape.Adjuster {
	for _, s := range s {
		diagram.applyStyle(s)
		diagram.Append(s)
	}
	return shape.NewAdjuster(s...)
}

// PlaceGrid place all the shapes into a grid starting at X,Y
// position. Row height is adapted to heighest element.
func (diagram *Diagram) PlaceGrid(cols, X, Y int, s ...shape.Shape) {
	row := make([]shape.Shape, cols)
	var x, y int
	var h shape.Shape
	for i, s := range s {
		switch {
		case i == 0:
			diagram.Place(s).At(X, Y)
		case y == 0:
			diagram.Place(s).RightOf(row[x-1])
			//			diagram.HAlignCenter(row[x-1], s)
		default:
			diagram.Place(s).Below(h)
			diagram.VAlignCenter(row[x], s)
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

// LinkAll places an arrow between the shapes, s0->s1->...->sn
func (diagram *Diagram) LinkAll(s ...shape.Shape) {
	for i, next := range s[1:] {
		diagram.Place(shape.NewArrowBetween(s[i], next))
	}
}

func (diagram *Diagram) Link(from, to shape.Shape, txt string) {
	lnk := shape.NewArrowBetween(from, to)
	diagram.Place(lnk)
	label := shape.NewLabel(txt)
	diagram.Place(label).Above(lnk, 20)
	diagram.VAlignCenter(lnk, label)
}

func (diagram *Diagram) applyStyle(s interface{}) {
	if s, ok := s.(shape.HasFont); ok {
		s.SetFont(diagram.Font)
	}
	if s, ok := s.(shape.HasTextPad); ok {
		s.SetTextPad(diagram.TextPad)
	}
}

// SaveAs saves the diagram to filename as SVG
func (d *Diagram) SaveAs(filename string) error {
	return saveAs(d, d.Style, filename)
}

func (d *Diagram) WriteSvg(w io.Writer) error {
	if d.Width == 0 && d.Height == 0 {
		d.AdaptSize()
	}
	if d.Caption != nil {
		margin := 30
		x := (d.Width - d.Caption.Width()) / 2
		if x < 0 {
			x = 0
		}
		d.Place(d.Caption).At(x, d.Height+margin)
		d.AdaptSize()
		d.Height += d.Caption.Font.Height / 2 // Fit protruding letters like 'g'
	}
	return d.Svg.WriteSvg(w)
}

// AdaptSize adapts the diagram size to the shapes inside it so all
// are visible. Returns the new width and height
func (diagram *Diagram) AdaptSize() (int, int) {
	for _, s := range diagram.Content {
		x, y := s.Position()
		switch s := s.(type) {
		case *shape.Line:
			x = min(s.Start.X, s.End.X)
			y = min(s.Start.Y, s.End.Y)
		case *shape.Arrow:
			x = min(s.Start.X, s.End.X)
			y = min(s.Start.Y, s.End.Y)
		}
		w := x + s.Width()
		if w > diagram.Width {
			diagram.Width = w
		}
		h := y + s.Height()
		if h > diagram.Height {
			diagram.Height = h
		}
	}
	return diagram.Width, diagram.Height
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SetHeight sets a fixed height in pixels.
func (d *Diagram) SetHeight(h int) {
	d.Height = h
}

// SetWidth sets a fixe width in pixels.
func (d *Diagram) SetWidth(w int) {
	d.Width = w
}

func (d *Diagram) SetCaption(txt string) {
	l := shape.NewLabel(txt)
	l.SetClass("caption")
	d.Caption = l
}
