package design

import (
	"io"
	"reflect"

	"github.com/gregoryv/draw/shape"
)

// NewClassDiagram returns a diagram representing structs and
// interfaces.  Relations are reflected from the types and drawn as
// arrows.
func NewClassDiagram() *ClassDiagram {
	return &ClassDiagram{
		Diagram:    NewDiagram(),
		interfaces: make([]VRecord, 0),
		structs:    make([]VRecord, 0),
		slices:     make([]VRecord, 0),
	}
}

type ClassDiagram struct {
	*Diagram

	interfaces []VRecord
	structs    []VRecord
	slices     []VRecord
}

func (d *ClassDiagram) Interface(obj interface{}) VRecord {
	vr := NewVRecord(obj)
	d.interfaces = append(d.interfaces, *vr)
	return *vr
}

func (d *ClassDiagram) Struct(obj interface{}) VRecord {
	vr := NewVRecord(obj)
	d.structs = append(d.structs, *vr)
	return *vr
}

func (d *ClassDiagram) Slice(obj interface{}) VRecord {
	vr := NewVRecord(obj)
	d.slices = append(d.slices, *vr)
	return *vr
}

// WriteSVG renders the diagram as SVG to the given writer.
func (d *ClassDiagram) WriteSVG(w io.Writer) error {
	rel := d.implements()
	rel = append(rel, d.compositions()...)
	for _, s := range rel {
		s, _ := s.(shape.Shape)
		d.Diagram.Prepend(s)
	}
	return d.Diagram.WriteSVG(w)
}

func (d *ClassDiagram) implements() []shape.Shape {
	rel := make([]shape.Shape, 0)
	for _, struct_ := range d.structs {
		for _, iface := range d.interfaces {
			if struct_.Implements(&iface) {
				arrow := shape.NewArrowBetween(struct_, iface)
				arrow.SetClass("implements-arrow")
				arrow.Head.SetClass("implements-arrow-head")
				rel = append(rel, arrow)
			}
		}
	}
	return rel
}

func (d *ClassDiagram) compositions() []shape.Shape {
	rel := make([]shape.Shape, 0)
	for _, struct_ := range d.structs {
		for _, struct2 := range d.structs {
			if struct_.ComposedOf(&struct2) {
				arrow := shape.NewArrowBetween(struct_, struct2)
				arrow.Tail = shape.NewDiamond()
				arrow.SetClass("compose-arrow")
				arrow.Tail.SetClass("compose-arrow-tail")
				rel = append(rel, arrow)
			}
			if struct_.Aggregates(&struct2) {
				arrow := shape.NewArrowBetween(struct_, struct2)
				arrow.Tail = shape.NewDiamond()
				arrow.SetClass("aggregate-arrow")
				arrow.Tail.SetClass("aggregate-arrow-tail")
				rel = append(rel, arrow)
			}
		}
		for _, slice := range d.slices {
			if struct_.ComposedOf(&slice) {
				arrow := shape.NewArrowBetween(struct_, slice)
				arrow.Tail = shape.NewDiamond()
				arrow.SetClass("compose-arrow")
				arrow.Tail.SetClass("compose-arrow-tail")
				rel = append(rel, arrow)
			}
			if slice.ComposedOf(&struct_) {
				arrow := shape.NewArrowBetween(slice, struct_)
				arrow.Tail = shape.NewDiamond()
				arrow.SetClass("compose-arrow")
				arrow.Tail.SetClass("compose-arrow-tail")
				rel = append(rel, arrow)
			}
			if struct_.Aggregates(&slice) {
				arrow := shape.NewArrowBetween(struct_, slice)
				arrow.Tail = shape.NewDiamond()
				arrow.SetClass("aggregate-arrow")
				arrow.Tail.SetClass("aggregate-arrow-tail")
				rel = append(rel, arrow)
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
				if field.Type == struct2.t || field.Type == reflect.PtrTo(struct2.t) {
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

// Inline returns rendered SVG with inlined style
func (d *ClassDiagram) Inline() string {
	return inline(d, d.Style)
}

// String returns rendered SVG
func (d *ClassDiagram) String() string { return toString(d) }

// Relation defines a relation between two records
type Relation struct {
	from, to *shape.Record
}
