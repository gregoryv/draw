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
		// always use pointer as the language works this way
		addMethods(rec, reflect.PtrTo(t))
	}
	if t.Kind() == reflect.Interface {
		addMethods(rec, t)
	}
	if t.Kind() == reflect.Slice {
		addMethods(rec, t)
	}
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
		r.Methods = append(r.Methods, m.Name+"()")
	}
}

// todo here so we can toggle manually added private methods
func isPublic(name string) bool {
	up := bytes.ToUpper([]byte(name))
	return []byte(name)[0] == up[0]
}

// VRecord represents a type struct or interface as a record shape.
type VRecord struct {
	*shape.Record
	t reflect.Type
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
		if field.Type == d.t || field.Type == reflect.SliceOf(d.t) {
			return true
		}
	}
	return false
}

func (vr *VRecord) Aggregates(d *VRecord) bool {
	for i := 0; i < vr.t.NumField(); i++ {
		field := vr.t.Field(i)
		if field.Type == reflect.PtrTo(d.t) || field.Type == reflect.SliceOf(reflect.PtrTo(d.t)) {
			return true
		}
	}
	return false
}
