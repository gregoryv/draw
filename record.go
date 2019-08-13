package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/go-design/svg"
	"github.com/gregoryv/go-design/xml"
)

func NewRecord(v interface{}) *Record {
	record := &Record{
		v:          reflect.TypeOf(v),
		StyleGuide: DefaultStyle,
	}
	return record
}

type Record struct {
	Pos
	v reflect.Type
	StyleGuide

	showPublicFields bool
}

func (record *Record) WriteTo(out io.Writer) (int, error) {
	all := make(Drawables, 0)
	x, y := record.X(), record.Y()
	w, h := record.Width(), record.Height()
	offset := record.Offset(x, y)
	padLeft := record.PaddingLeft
	name := record.v.Name()

	attributes := xml.Attributes{class("record")}
	if record.HasSpecialStyle() {
		attributes = append(attributes, record.FillStroke())
	}
	all = append(all,
		svg.Rect(x, y, w, h, attributes...),
		svg.Text(x+padLeft, offset.Line(1)-record.PaddingTop, name),
	)

	// Public fields
	if record.showPublicFields {
		attributes := xml.Attributes{}
		if record.HasSpecialStyle() {
			attributes = append(attributes, record.Stroke())
		}
		all = append(all,
			svg.Line(
				x, offset.Line(1)+record.PaddingBottom,
				x+w, offset.Line(1)+record.PaddingBottom,
				attributes...,
			),
		)
		for i := 0; i < record.v.NumField(); i++ {
			yOffset := record.PaddingTop + record.StyleGuide.Height(i+2)
			field := record.v.Field(i)
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

func (record *Record) Center() (x int, y int) {
	return record.X() + record.Width()/2, record.Y() + record.Height()/2

}

func (record *Record) Width() int {
	n := widthOf(record.v.Name())
	if record.showPublicFields {
		for i := 0; i < record.v.NumField(); i++ {
			field := record.v.Field(i)
			w := widthOf(field.Name)
			if w > n {
				n = w
			}
		}
	}
	return n
}

func (record *Record) Height() int {
	n := 1
	if record.showPublicFields {
		n += record.v.NumField()
		n += 1 // separators
	}
	return DefaultStyle.Height(n)
}

func (record *Record) ShowFields() { record.showPublicFields = true }

func class(v string) xml.Attribute { return attr("class", v) }
func style(v string) xml.Attribute { return attr("style", v) }
func attr(key, val string) xml.Attribute {
	return xml.NewAttribute(key, val)
}

func (a *Record) AreLinked(b *Record) bool {
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

func (record *Record) WithFields() *Record {
	record.showPublicFields = true
	return record
}
