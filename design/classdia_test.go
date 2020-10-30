package design

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestClassDiagram_OtherType(t *testing.T) {
	newClassDiagram().SaveAs("testdata/classdia_test.svg")
	data := newClassDiagram().String()
	ioutil.WriteFile("testdata/classdia_inlined.svg", []byte(data), 0644)
}

func TestClassDiagram_Inline(t *testing.T) {
	got := newClassDiagram().Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
}

func TestClassDiagram_String(t *testing.T) {
	got := newClassDiagram().String()
	if !strings.Contains(got, "class") {
		t.Error("missing class attributes\n", got)
	}
}

func newClassDiagram() *ClassDiagram {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Slice(make(Wheels, 4))
		stringer = d.Interface((*fmt.Stringer)(nil))
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)
	return d
}

type Car struct {
	Model string
	Wheels
}

type Wheel struct {
	Make string
}

type Wheels []Wheel

// String
func (me *Wheel) String() string { return me.Make }
