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

// NewClassDiagram returns a diagram representing structs and
// interfaces.  Relations are reflected from the types and drawn as
// arrows.
func NewClassDiagram() *ClassDiagram {
	return &ClassDiagram{
		Diagram:    NewDiagram(),
		Interfaces: make([]VRecord, 0),
		Structs:    make([]VRecord, 0),
	}
}

// WriteSvg renders the diagram as SVG to the given writer.
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
					// TODO
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

// Place places adds the record to the diagram returning an adjuster
// for positioning.
func (d *ClassDiagram) Place(vr VRecord) *shape.Adjuster {
	if vr.isStruct {
		d.Structs = append(d.Structs, vr)
	} else {
		d.Interfaces = append(d.Interfaces, vr)
	}
	return d.Diagram.Place(vr.Record)
}

// Relation defines a relation between two records
type Relation struct {
	from, to *shape.Record
}

// NewStruct returns a VRecord of the given object, panics if not
// struct.
func NewStruct(obj interface{}) VRecord {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Expected struct kind got %v", t.Kind()))
	}
	return VRecord{
		Record:   shape.NewStructRecord(obj),
		t:        t,
		isStruct: true,
	}
}

// VRecord represents a type struct or interface as a record shape.
type VRecord struct {
	*shape.Record
	t        reflect.Type
	isStruct bool
}

func (vr *VRecord) TitleOnly() *VRecord {
	vr.HideFields()
	vr.HideMethods()
	return vr
}

// NewInterface returns a VRecord of the given object, panics if not
// interface.
func NewInterface(obj interface{}) VRecord {
	t := reflect.TypeOf(obj).Elem()
	if t.Kind() != reflect.Interface {
		panic(fmt.Sprintf("Expected ptr kind got %v", t.Kind()))
	}
	return VRecord{
		Record:   shape.NewInterfaceRecord(obj),
		t:        t,
		isStruct: false,
	}
}

// SaveAs saves the diagram to filename as SVG
func (d *ClassDiagram) SaveAs(filename string) error {
	return saveAs(d, filename)
}
