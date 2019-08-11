package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/go-design/xml"
)

func NewComponent(v interface{}) *Component {
	comp := &Component{
		Label: reflect.TypeOf(v).Name(),
		v:     v,
	}
	return comp
}

type Component struct {
	Label string
	v     interface{}
}

func (comp *Component) WriteTo(w io.Writer) (int, error) {
	all := make(Drawables, 0)
	xp, yp := 30, 20
	all = append(all,
		xml.NewNode(Element_rect, x(xp), y(yp),
			width(toFit(comp.Label)), height("150"),
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
		xml.NewNode(Element_rect, x(xp-5), y(yp+5), width("10"), height("10"),
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
	)

	text := xml.NewNode(Element_text, x(xp+20), y(yp+25), fill("black"))
	text.Append(xml.CData(comp.Label))
	all = append(all, text)
	return all.WriteTo(w)
}

func (comp *Component) Style() *StyleGuide {
	return DefaultStyle
}
