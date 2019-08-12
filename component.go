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
	Pos
	v reflect.Type

	showPublicFields bool
}

func (comp *Component) Copy() *Component {
	return &Component{
		Pos: comp.Pos,
		v:   comp.v,
	}
}

func (comp *Component) WriteTo(out io.Writer) (int, error) {
	all := make(Drawables, 0)
	x, y := comp.X(), comp.Y()
	w, h := comp.Width(), comp.Height()
	s := DefaultStyle
	offset := s.Offset(x, y)
	padLeft := s.PaddingLeft
	name := comp.v.Name()
	all = append(all,
		svg.Rect(x, y, w, h, class("component")),
		//svg.Rect(x-5, y+5, 10, 10, class("smallbox")),
		// Name of type
		svg.Text(x+padLeft, offset.Line(1)-s.PaddingTop, name),
	)

	// Public fields
	if comp.showPublicFields {
		all = append(all,
			svg.Line(
				x, offset.Line(1)+s.PaddingBottom,
				x+w, offset.Line(1)+s.PaddingBottom,
			),
		)
		for i := 0; i < comp.v.NumField(); i++ {
			yOffset := s.PaddingTop + s.Height(i+2)
			field := comp.v.Field(i)
			all = append(all,
				svg.Text(x+padLeft, y+yOffset, field.Name),
			)
		}
	}
	return all.WriteTo(out)
}

func nameAndType(field reflect.StructField) string {
	typ := field.Type.Kind().String()
	return field.Name + " " + typ
}

func (comp *Component) Center() (x int, y int) {
	return comp.X() + comp.Width()/2, comp.Y() + comp.Height()/2

}

func (comp *Component) Width() int {
	n := widthOf(comp.v.Name())
	for i := 0; i < comp.v.NumField(); i++ {
		field := comp.v.Field(i)
		w := widthOf(field.Name)
		if w > n {
			n = w
		}
	}
	return n
}

func (comp *Component) Height() int {
	n := 1
	if comp.showPublicFields {
		n += comp.v.NumField()
		n += 1 // separators
	}
	return DefaultStyle.Height(n)
}

func (comp *Component) Style() *StyleGuide { return DefaultStyle }
func (comp *Component) ShowFields()        { comp.showPublicFields = true }

func class(v string) xml.Attribute { return attr("class", v) }
func attr(key, val string) xml.Attribute {
	return xml.NewAttribute(key, val)
}

func (a *Component) AreLinked(b *Component) bool {
	for i := 0; i < a.v.NumField(); i++ {
		if linked(a.v.Field(i).Type, b.v) {
			return true
		}
	}
	for i := 0; i < b.v.NumField(); i++ {
		if linked(b.v.Field(i).Type, a.v) {
			return true
		}
	}
	return false
}

func linked(from, to reflect.Type) bool {
	return from == to || from == reflect.PtrTo(to)
}

func (comp *Component) WithFields() *Component {
	comp.showPublicFields = true
	return comp
}
