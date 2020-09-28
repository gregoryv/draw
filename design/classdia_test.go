package design

import (
	"fmt"
	"testing"
)

func TestClassDiagram_OtherType(t *testing.T) {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Struct(Wheel{})
		stringer = d.Interface((*fmt.Stringer)(nil))
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)

	d.SaveAs("testdata/classdia_test.svg")
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
