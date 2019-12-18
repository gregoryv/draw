package design

import (
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestVRecord(t *testing.T) {
	r := NewStruct(VRecord{})
	before := len(r.Fields)
	r.TitleOnly()
	got := len(r.Fields)
	assert := asserter.New(t)
	assert(got != before).Error("Did not hide fields")
}

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

type C struct{}

func TestVRecord_ComposedOf(t *testing.T) {
	ok := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if !A.ComposedOf(&B) {
			t.Fail()
		}
	}
	ok(struct{ c C }{}, C{})

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewStruct(a)
		B := NewStruct(b)
		if A.ComposedOf(&B) {
			t.Fail()
		}
	}
	bad(struct{ c *C }{}, C{})
}
