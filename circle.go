package design

import (
	"bytes"
	"io"
	"text/template"
)

type Circle struct {
	Pos
	StyleGuide
	Label Label
}

func NewCircle(label string) *Circle {
	return &Circle{
		StyleGuide: DefaultStyle,
		Label: Label{
			StyleGuide: DefaultStyle,
			Text:       label,
		},
	}
}

func (circle *Circle) Height() int   { return circle.Diameter() }
func (circle *Circle) Width() int    { return circle.Diameter() }
func (circle *Circle) Diameter() int { return widthOf(circle.Label.Text) }
func (circle *Circle) Radius() int   { return circle.Diameter() / 2 }

// todo SetX should affect label

func (circle *Circle) WriteTo(w io.Writer) (int, error) {
	xml := `<circle cx="{{.X}}" cy="{{.Y}}" r="{{.Radius}}" {{.FillStrokeS}}/>
{{.Label}}
`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, circle)
	n, err := buf.WriteTo(w)
	return int(n), err // todo switch to int6r
}
