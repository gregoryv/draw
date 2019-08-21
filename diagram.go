package design

import (
	"io"
	"reflect"

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
	Font    shape.Font
	TextPad shape.Padding
	Pad     shape.Padding
}

func (diagram *Diagram) Place(obj interface{}) *shape.Adjuster {
	rec := reflectRecord(obj)
	rec.Font = diagram.Font
	rec.Pad = diagram.TextPad
	diagram.Append(rec)
	return shape.NewAdjuster(rec)
}

func reflectRecord(obj interface{}) *shape.Record {
	t := reflect.TypeOf(obj)
	rec := shape.NewRecord(t.Name())
	return rec
}

func (diagram *Diagram) SaveAs(filename string) error {
	return saveAs(diagram, filename)
}

func (diagram *Diagram) WriteSvg(w io.Writer) error {
	diagram.AdaptSize()
	return diagram.Svg.WriteSvg(w)
}

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
