package design

import (
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestNewStruct(t *testing.T) {
	x := struct {
		Field string
	}{}
	s := NewStruct(x)
	if len(s.Fields) != 1 {
		t.Error("Expected one field")
	}

	defer mustCatchPanic(t)
	NewStruct((*io.Writer)(nil))
}

func TestNewInterface(t *testing.T) {
	i := NewInterface((*io.Writer)(nil))
	if len(i.Methods) != 1 {
		t.Error("Expected one method")
	}

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
