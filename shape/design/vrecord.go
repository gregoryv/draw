package design

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/gregoryv/draw/shape"
)

func NewVRecord(v interface{}) *VRecord {
	t := reflect.TypeOf(v)
	title := fmt.Sprintf("%s %s", t, t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Interface {
			title = fmt.Sprintf("%s %s", t, t.Kind())
		}
	}
	rec := shape.NewRecord(title)

	if t.Kind() == reflect.Struct {
		addFields(rec, t)
		addMethods(rec, t)
	}
	if t.Kind() == reflect.Interface {
		addMethods(rec, t)
	}

	// todo add methods and fields if any
	return &VRecord{
		Record: rec,
		t:      t,
	}
}

func addFields(r *shape.Record, t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if isPublic(field.Name) {
			r.Fields = append(r.Fields, field.Name)
		}
	}
}

func addMethods(r *shape.Record, t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if isPublic(m.Name) {
			r.Methods = append(r.Methods, m.Name+"()")
		}
	}
}

func isPublic(name string) bool {
	up := bytes.ToUpper([]byte(name))
	return []byte(name)[0] == up[0]
}

// NewStructRecord returns a record shape based on a Go struct type.
// Reflection is used.
func NewStructRecord(obj interface{}) *shape.Record {
	t := reflect.TypeOf(obj)
	rec := shape.NewRecord(t.String() + " struct")
	addFields(rec, t)
	addMethods(rec, reflect.PtrTo(t))
	return rec
}

func NewInterfaceRecord(obj interface{}) *shape.Record {
	t := reflect.TypeOf(obj).Elem()
	rec := shape.NewRecord(t.String() + " interface")
	addMethods(rec, t)
	return rec
}

// VRecord represents a type struct or interface as a record shape.
type VRecord struct {
	*shape.Record
	t reflect.Type
}

// NewInterface returns a VRecord of the given object, panics if not
// interface.
func NewInterface(obj interface{}) VRecord {
	r := NewVRecord(obj)
	return *r
}

// TitleOnly hides fields and methods.
func (vr *VRecord) TitleOnly() {
	vr.HideFields()
	vr.HideMethods()
}

func (vr *VRecord) Implements(iface *VRecord) bool {
	return reflect.PtrTo(vr.t).Implements(iface.t)
}

func (vr *VRecord) ComposedOf(d *VRecord) bool {
	for i := 0; i < vr.t.NumField(); i++ {
		field := vr.t.Field(i)
		if field.Type == d.t {
			return true
		}
	}
	return false
}

func (vr *VRecord) Aggregates(d *VRecord) bool {
	for i := 0; i < vr.t.NumField(); i++ {
		field := vr.t.Field(i)
		if field.Type == reflect.PtrTo(d.t) {
			return true
		}
	}
	return false
}
