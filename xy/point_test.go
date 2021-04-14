package xy

import (
	"fmt"
	"testing"
)

func TestOnePoint(t *testing.T) {
	it := OnePoint{t, &Point{10, 10}}
	q := Point{10, 10}
	it.IsNotLeftOf(q)
	it.IsNotRightOf(q)
	it.IsNotAbove(q)
	it.IsNotBelow(q)
	it.IsEqualTo(q)
	it.HasQuickAccessToCoordinates()
	it.ImplementsStringer()
}

type OnePoint struct {
	*testing.T
	*Point
}

func (t OnePoint) IsLeftOf(q Point) {
	t.Helper()
	if !t.LeftOf(q) {
		t.Errorf("%v should be left of %v", t.Point, q)
	}
}

func (t OnePoint) IsNotLeftOf(q Point) {
	t.Helper()
	if t.LeftOf(q) {
		t.Errorf("%v should not be left of %v", t.Point, q)
	}
}

func (t OnePoint) IsNotRightOf(q Point) {
	t.Helper()
	if t.RightOf(q) {
		t.Errorf("%v should not be right of %v", t.Point, q)
	}
}

func (t OnePoint) IsNotAbove(q Point) {
	t.Helper()
	if t.Above(q) {
		t.Errorf("%v should not be above %v", t.Point, q)
	}
}

func (t OnePoint) IsNotBelow(q Point) {
	t.Helper()
	if t.Below(q) {
		t.Errorf("%v should not be below %v", t.Point, q)
	}
}

func (t OnePoint) IsEqualTo(q Point) {
	t.Helper()
	if !t.Equals(q) {
		t.Errorf("%v should be equal to %v", t.Point, q)
	}
}

func (t OnePoint) HasQuickAccessToCoordinates() {
	t.Helper()
	x, y := t.XY()
	if x != t.X || y != t.Y {
		t.Errorf("Coordinates do not match %v %s", t.Point, t.String())
	}
}

func (t OnePoint) ImplementsStringer() {
	_ = fmt.Stringer(t.Point)
}

func TestPoint_Distance(t *testing.T) {
	var p Point
	cases := []struct {
		q   Point
		exp float64
	}{
		{Point{0, 2}, 2},
		{Point{2, 0}, 2},
		{Point{-2, 0}, 2},
		{Point{0, -2}, 2},
	}
	for _, c := range cases {
		got := p.Distance(c.q)
		if got != c.exp {
			t.Error(got, c.exp)
		}
	}
}
