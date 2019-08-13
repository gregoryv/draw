package design

import (
	"io"

	"github.com/gregoryv/go-design/svg"
	"github.com/gregoryv/go-design/xml"
)

type Circle struct {
	Pos
	StyleGuide
	Label string
}

func NewCircle(label string) *Circle {
	return &Circle{
		StyleGuide: DefaultStyle,
		Label:      label,
	}
}

func (circle *Circle) Height() int   { return circle.Diameter() }
func (circle *Circle) Width() int    { return circle.Diameter() }
func (circle *Circle) Diameter() int { return widthOf(circle.Label) }

func (circle *Circle) WriteTo(w io.Writer) (int, error) {
	radius := circle.Diameter() / 2
	attributes := xml.Attributes{}
	if circle.HasSpecialStyle() {
		attributes = append(attributes, circle.FillStroke())
	}
	all := make(Drawables, 2)
	all[0] = svg.Circle(circle.X(), circle.Y(), radius, attributes...)
	textX := circle.X() - radius + circle.PaddingLeft
	all[1] = svg.Text(textX, circle.Y(), circle.Label,
		attr("font-family", circle.FontFamily),
		circle.FillStroke(),
	)

	return all.WriteTo(w)
}
