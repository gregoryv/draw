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
