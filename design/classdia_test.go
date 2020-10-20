package design

import (
	"fmt"
	"strings"
	"testing"
)

func TestClassDiagram_OtherType(t *testing.T) {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Struct(&Wheel{})
		stringer = d.Interface((*fmt.Stringer)(nil))
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)

	d.SaveAs("testdata/classdia_test.svg")
}

func TestClassDiagram_Inline(t *testing.T) {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Struct(&Wheel{})
		stringer = d.Interface((*fmt.Stringer)(nil))
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)

	got := d.Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
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
