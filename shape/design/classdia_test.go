package design

import (
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestDesignPackage(t *testing.T) {
	it := NewDesignPackage(t)
	it.CanCreateRecordFromStruct()
	it.CanCreateRecordFromInterface()
}

func NewDesignPackage(t *testing.T) *DesignPackage {
	return &DesignPackage{T: t}
}

type DesignPackage struct {
	*testing.T
}

func (t *DesignPackage) CanCreateRecordFromStruct() {
	t.Log("Can create record from struct")
	x := struct {
		Field string
	}{}
	NewStruct(x)

	defer mustCatchPanic(t)
	NewStruct((*io.Writer)(nil))
}

func (t *DesignPackage) CanCreateRecordFromInterface() {
	t.Log("Can create record from interface")
	NewInterface((*io.Writer)(nil))
	defer mustCatchPanic(t)
	NewInterface(t)
}

func mustCatchPanic(t asserter.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("should panic")
	}
}
