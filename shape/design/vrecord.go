package design

import (
	"fmt"
	"reflect"

	"github.com/gregoryv/draw/shape"
)

// VRecord represents a type struct or interface as a record shape.
type VRecord struct {
	*shape.Record
	t        reflect.Type
	isStruct bool
}

func (vr *VRecord) TitleOnly() {
	vr.HideFields()
	vr.HideMethods()
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
