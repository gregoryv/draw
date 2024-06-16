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

type (
	myOwn int
	myStr struct{ f string }
)

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
	ok := func(a, b interface{}, msg string) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if !A.Aggregates(B) {
			t.Errorf("%s\n%v aggregates of %v", msg, A, B)
		}
	}

	// an aggregate is defined as a relation between two types which
	// may be nil at runtime.
	ok(struct{ c *C }{}, C{}, "pointer to struct")
	ok(struct{ c []*C }{}, C{}, "slice of pointers to struct")
	ok(struct{ c []C }{}, C{}, "slice of structs")
	ok(struct{ c *MySlice }{}, MySlice{}, "named slice")
	ok(struct{ c fmt.Stringer }{}, (*fmt.Stringer)(nil), "interface field")
	ok(struct{ c map[int]string }{}, map[int]string{}, "a map")
	ok(struct{ c chan int }{}, make(chan int, 0), "chan")
	ok(struct{ c func() }{}, func() {}, "func")
	ok(struct{ c F }{}, F(func() {}), "func")

	bad := func(a, b interface{}) {
		t.Helper()
		A := NewVRecord(a)
		B := NewVRecord(b)
		if A.Aggregates(B) {
			t.Errorf("%v doesn't aggregate %v", A, B)
		}
	}
	bad(struct{ c C }{}, C{})
	bad(1, C{})
}

func noop() {}

type (
	F       func()
	C       struct{}
	MySlice []C
)
