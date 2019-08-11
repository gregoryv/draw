package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/go-design/svg"
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
		svg.Rect(xp, yp, comp.Width(), h,
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
		svg.Rect(xp-5, yp+5, 10, 10,
			style("fill:#ffffcc;stroke:black;stroke-width:1"),
		),
	)
	// Name of type
	all = append(all, svg.Text(xp+20, yp+25, name, fill("black")))

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

func style(v string) xml.Attribute { return attr("style", v) }
func fill(v string) xml.Attribute  { return attr("fill", v) }
func attr(key, val string) xml.Attribute {
	return xml.NewAttribute(key, val)
}
