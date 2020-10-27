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

func TestClassDiagram_Inline_same_as_stringer(t *testing.T) {
	got := newClassDiagram().Inline()
	exp := newClassDiagram().String()
	if got != exp {
		t.Errorf(
			"inline is \n%s\n\n and stringer is \n %s",
			got, exp,
		)
	}
}

func newClassDiagram() *ClassDiagram {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Struct(&Wheel{})
		stringer = d.Interface((*fmt.Stringer)(nil))
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)
	return d
}

type Car struct {
	Model  string
	Wheels []*Wheel
}

type Wheel struct {
	Make string
}

// String
func (me *Wheel) String() string { return me.Make }
