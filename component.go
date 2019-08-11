package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/go-design/xml"
)

func NewComponent(v interface{}) *Component {
	comp := &Component{
		v: reflect.TypeOf(v),
	}
	return comp
}

type Component struct {
	v reflect.Type

	showPublicFields bool
}

func (comp *Component) WriteTo(w io.Writer) (int, error) {
	all := make(Drawables, 0)
	xp, yp := 30, 20
	h := 150
	name := comp.v.Name()
	all = append(all,
		Rect(xp, yp, comp.Width(), h,
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
		xml.NewNode(Element_rect, x(xp-5), y(yp+5), width(10), height(10),
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
	)
	// Name of type
	text := xml.NewNode(Element_text, x(xp+20), y(yp+25), fill("black"))
	text.Append(xml.CData(name))
	all = append(all, text)

	// Line separator

	// Public fields
	return all.WriteTo(w)
}

func (comp *Component) Width() int {
	// todo find widest
	return widthOf(comp.v.Name())
}
func (comp *Component) Style() *StyleGuide { return DefaultStyle }
func (comp *Component) ShowPublicFields()  { comp.showPublicFields = true }
