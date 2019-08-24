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
				// k won't work for all cases
				k := 1
				if iface.X > struct_.X {
					k = 0
				}
				line := &shape.Arrow{
					X1: struct_.X + struct_.Width()/2,
					Y1: struct_.Y + struct_.Height()/2,
					X2: iface.X + k*iface.Width(),
					Y2: iface.Y + iface.Height()/2,
				}
				rel = append(rel, line)
			}
		}
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
