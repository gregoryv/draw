package design

import (
	"io"

	"github.com/gregoryv/go-design/shape"
)

func NewDiagram() Diagram {
	return Diagram{
		Font:    shape.Font{Height: 9, Width: 7, LineHeight: 15},
		TextPad: shape.Padding{Left: 10, Top: 2, Bottom: 7, Right: 10},
		Pad:     shape.Padding{Left: 10, Top: 20, Bottom: 7, Right: 10},
	}
}

type Diagram struct {
	shape.Svg

	width, height int

	Font    shape.Font
	TextPad shape.Padding
	Pad     shape.Padding
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
	diagram.AdaptSize()
	return diagram.Svg.WriteSvg(w)
}

// AdaptSize adapts the diagram size to the shapes inside it so all
// are visible.
func (diagram *Diagram) AdaptSize() {
	width := diagram.Width
	height := diagram.Height
	if width > 0 && height > 0 {
		return
	}
	for _, s := range diagram.Content {
		x, y := s.Position()
		w := x + s.Width()
		if w > width {
			width = w
		}
		h := y + s.Height()
		if h > height {
			height = h
		}
	}
	diagram.Width = width
	diagram.Height = height
}

func (d *Diagram) SetHeight(h int) { d.height = h }
func (d *Diagram) SetWidth(w int)  { d.width = w }

func (*Diagram) HAlignCenter(shapes ...shape.Shape) {
	shape.HAlign(shape.Center, shapes...)
}

func (*Diagram) HAlignTop(shapes ...shape.Shape) {
	shape.HAlign(shape.Top, shapes...)
}

func (*Diagram) HAlignBottom(shapes ...shape.Shape) {
	shape.HAlign(shape.Bottom, shapes...)
}

func (*Diagram) VAlignCenter(shapes ...shape.Shape) {
	shape.VAlign(shape.Center, shapes...)
}

func (*Diagram) VAlignLeft(shapes ...shape.Shape) {
	shape.VAlign(shape.Left, shapes...)
}

func (*Diagram) VAlignRight(shapes ...shape.Shape) {
	shape.VAlign(shape.Right, shapes...)
}
