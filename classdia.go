package design

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gregoryv/go-design/shape"
)

type ClassDiagram struct {
	Diagram

	interfaces []VRecord
	structs    []VRecord
}

// NewClassDiagram returns a diagram representing structs and
// interfaces.  Relations are reflected from the types and drawn as
// arrows.
func NewClassDiagram() *ClassDiagram {
	return &ClassDiagram{
		Diagram:    NewDiagram(),
		interfaces: make([]VRecord, 0),
		structs:    make([]VRecord, 0),
	}
}

func (d *ClassDiagram) Interface(obj interface{}) VRecord {
	vr := NewInterface(obj)
	d.interfaces = append(d.interfaces, vr)
	return vr
}

func (d *ClassDiagram) Struct(obj interface{}) VRecord {
	vr := NewStruct(obj)
	d.structs = append(d.structs, vr)
	return vr
}

// WriteSvg renders the diagram as SVG to the given writer.
func (d *ClassDiagram) WriteSvg(w io.Writer) error {
	rel := d.implements()
	rel = append(rel, d.composes()...)
	d.Diagram.Prepend(rel...)
	return d.Diagram.WriteSvg(w)
}

func (d *ClassDiagram) implements() []shape.Shape {
	rel := make([]shape.Shape, 0)
	for _, struct_ := range d.structs {
		for _, iface := range d.interfaces {
			if reflect.PtrTo(struct_.t).Implements(iface.t) {
				arrow := shape.NewArrowBetween(struct_, iface)
				arrow.SetClass("implements-arrow")
				arrow.Head.SetClass("implements-arrow-head")
				rel = append(rel, arrow)
			}
		}
	}
	return rel
}

func (d *ClassDiagram) composes() []shape.Shape {
	rel := make([]shape.Shape, 0)
	for _, struct_ := range d.structs {
		for i := 0; i < struct_.t.NumField(); i++ {
			field := struct_.t.Field(i)
			for _, struct2 := range d.structs {
				if field.Type == struct2.t {
					// todo use composition tail shape
					arrow := shape.NewArrowBetween(struct_, struct2)
					rel = append(rel, arrow)
				}
			}
		}
	}
	return rel
}

// HideRealizations hides all methods of structs that implement a
// visible interface.
func (d *ClassDiagram) HideRealizations() {
	for _, struct_ := range d.structs {
		for _, iface := range d.interfaces {
			if reflect.PtrTo(struct_.t).Implements(iface.t) {
				// Hide interface methods as they are visible
				// in the diagram already
				for _, m := range iface.Methods {
					struct_.HideMethod(m)
				}
			}
		}
	}
	for _, struct_ := range d.structs {
		for i := 0; i < struct_.t.NumField(); i++ {
			field := struct_.t.Field(i)
			for _, struct2 := range d.structs {
				if field.Type == struct2.t {
					for _, m := range struct2.Methods {
						struct_.HideMethod(m)
					}
				}
			}
		}
	}
}

// SaveAs saves the diagram to filename as SVG
func (d *ClassDiagram) SaveAs(filename string) error {
	return saveAs(d, d.Style, filename)
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
