package design

import (
	"fmt"
	"io"
	"testing"

	"github.com/gregoryv/asserter"
)

func Example_VRecord_methods() {
	vr := NewDetailedVRecord(&car{})
	for _, m := range vr.Record.Methods {
		fmt.Println(m)
	}
	// output:
	// Model(int) string
}

type car struct{}

func (me *car) Model(x int) string {
	return ""
}

type myOwn int
type myStr struct{ f string }

func Test_NewVRecord_types(t *testing.T) {
	ok := func(v interface{}, exp string) {
		vr := NewVRecord(v)
		got := vr.Title
		if got != exp {
			t.Error("got: ", got, "exp: ", exp)
		}
	}
	ok(myOwn(1), "design.myOwn int")
	ok(myStr{}, "design.myStr struct")
	ok((*io.Reader)(nil), "io.Reader interface")
}

func TestVRecord_TitleOnly_hides_fields(t *testing.T) {
	r := NewVRecord(VRecord{})
	before := len(r.Fields)
	r.TitleOnly()
	got := len(r.Fields)
	assert := asserter.New(t)
	assert(got != before).Fail()
}

func mustCatchPanic(t asserter.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("should panic")
	}
}

func TestVRecord_ComposedOf(t *testing.T) {
	ok := func(a, b interface{}) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if !A.ComposedOf(B) {
			t.Errorf("%v composes %v", A, B)
		}
	}
	ok(struct{ c C }{}, C{})

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if A.ComposedOf(B) {
			t.Errorf("%v doesn't compose %v", A, B)
		}
	}
	bad(struct{ c *C }{}, C{})
}

func TestVRecord_Aggregates(t *testing.T) {
	ok := func(a, b interface{}) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if !A.Aggregates(B) {
			t.Errorf("%v aggregates of %v", A, B)
		}
	}
	ok(struct{ c *C }{}, C{})
	ok(struct{ c []*C }{}, C{})
	ok(struct{ c *MySlice }{}, MySlice{})

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if A.Aggregates(B) {
			t.Errorf("%v doesn't aggregate %v", A, B)
		}
	}
	bad(struct{ c C }{}, C{})
}

type C struct{}
type MySlice []C
