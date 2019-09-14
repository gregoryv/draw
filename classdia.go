package design

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gregoryv/go-design/shape"
)

type ClassDiagram struct {
	Diagram

	Interfaces []VRecord
	Structs    []VRecord
}

func NewClassDiagram() *ClassDiagram {
	return &ClassDiagram{
		Diagram:    NewDiagram(),
		Interfaces: make([]VRecord, 0),
		Structs:    make([]VRecord, 0),
	}
}

func (d *ClassDiagram) WriteSvg(w io.Writer) error {
	rel := make([]shape.SvgWriterShape, 0)
	for _, struct_ := range d.Structs {
		for _, iface := range d.Interfaces {
			if reflect.PtrTo(struct_.t).Implements(iface.t) {
				// todo arrow is hidden by destination, calculate edge x,y
				arrow := shape.NewArrow(
					struct_.X+struct_.Width()/2,
					struct_.Y+struct_.Height()/2,
					iface.X+iface.Width()/2,
					iface.Y+iface.Height()/2,
				)
				switch {
				case arrow.DirQ1(), arrow.DirQ4():
					arrow.End.X -= iface.Width() / 2
				case arrow.DirQ2(), arrow.DirQ3():
					arrow.End.X += iface.Width() / 2
				}
				rel = append(rel, arrow)
				// todo, add implements label
			}
		}
		// todo, composition
	}
	d.Diagram.Prepend(rel...)
	return d.Diagram.WriteSvg(w)
}

func (d *ClassDiagram) Place(vr VRecord) *shape.Adjuster {
	if vr.isStruct {
		d.Structs = append(d.Structs, vr)
	} else {
		d.Interfaces = append(d.Interfaces, vr)
	}
	return d.Diagram.Place(vr.Record)
}

type Relation struct {
	from, to *shape.Record
}

type VRecord struct {
	*shape.Record
	t        reflect.Type
	isStruct bool
}

func NewStruct(obj interface{}) VRecord {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Expected struct kind got %v", t.Kind()))
	}
	return VRecord{shape.NewStructRecord(obj), t, true}
}

func NewInterface(obj interface{}) VRecord {
	t := reflect.TypeOf(obj).Elem()
	if t.Kind() != reflect.Interface {
		panic(fmt.Sprintf("Expected ptr kind got %v", t.Kind()))
	}
	return VRecord{shape.NewInterfaceRecord(obj), t, false}
}

// SaveAs saves the diagram to filename as SVG
func (d *ClassDiagram) SaveAs(filename string) error {
	return saveAs(d, filename)
}
