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
		Font:    shape.DefaultFont,
		TextPad: shape.DefaultTextPad,
		Pad:     shape.DefaultPad,
	}
}

// Diagram is a generic SVG image with box related styling
type Diagram struct {
	shape.Svg
	shape.Aligner

	Font    shape.Font    // Used to calculate width
	TextPad shape.Padding // Surrounding text
	Pad     shape.Padding // E.g. records
}

// Place adds the shape to the diagram returning an adjuster for
// positioning.
func (diagram *Diagram) Place(s shape.SvgWriterShape) *shape.Adjuster {
	diagram.applyStyle(s)
	diagram.Append(s)
	return shape.NewAdjuster(s)
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
func (diagram *Diagram) SaveAs(filename string) error {
	return saveAs(diagram, filename)
}

func (diagram *Diagram) WriteSvg(w io.Writer) error {
	if diagram.Width == 0 && diagram.Height == 0 {
		diagram.AdaptSize()
	}
	return diagram.Svg.WriteSvg(w)
}

// AdaptSize adapts the diagram size to the shapes inside it so all
// are visible. Returns the new width and height
func (diagram *Diagram) AdaptSize() (int, int) {
	for _, s := range diagram.Content {
		x, y := s.Position()
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

// SetHeight sets a fixed height in pixels.
func (d *Diagram) SetHeight(h int) {
	d.Height = h
}

// SetWidth sets a fixe width in pixels.
func (d *Diagram) SetWidth(w int) {
	d.Width = w
}
