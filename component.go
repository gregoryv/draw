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

	x, y int

	showPublicFields bool
}

func (comp *Component) WriteTo(out io.Writer) (int, error) {
	all := make(Drawables, 0)
	x, y := comp.x, comp.y
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
			all = append(all,
				svg.Text(x+padLeft, y+yOffset, comp.v.Field(i).Name),
			)
		}
	}
	return all.WriteTo(out)
}

func (comp *Component) Width() int {
	// todo find widest
	return widthOf(comp.v.Name())
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

func style(v string) xml.Attribute { return attr("style", v) }
func class(v string) xml.Attribute { return attr("class", v) }
func fill(v string) xml.Attribute  { return attr("fill", v) }
func attr(key, val string) xml.Attribute {
	return xml.NewAttribute(key, val)
}
