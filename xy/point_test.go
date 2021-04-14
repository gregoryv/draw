package xy

import (
	"testing"
)

func TestPoint_Distance(t *testing.T) {
	var p Point
	cases := []struct {
		q   Point
		exp float64
	}{
		{Point{0, 2}, 2},
	}
	for _, c := range cases {
		got := p.Distance(c.q)
		if got != c.exp {
			t.Error(got, c.exp)
		}
	}

}

func Test_OnePoint(t *testing.T) {
	it := OnePoint{t, Point{10, 10}}
	q := Point{10, 10}
	it.IsNotLeftOf(q)
	it.IsNotRightOf(q)
	it.IsNotAbove(q)
	it.is_not_below(q)
	it.is_same_as(q)
	it.has_quick_access_to_coordinates()
	it.is_stringable()
}

type OnePoint struct {
	*testing.T
	Point
}

func (t OnePoint) is_left_of(q Point) {
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

func (t OnePoint) is_not_below(q Point) {
	t.Helper()
	if t.Below(q) {
		t.Errorf("%v should not be below %v", t.Point, q)
	}
}

func (t OnePoint) is_same_as(q Point) {
	t.Helper()
	if !t.Equals(q) {
		t.Errorf("%v should be equal to %v", t.Point, q)
	}
}

func (t OnePoint) has_quick_access_to_coordinates() {
	t.Helper()
	x, y := t.XY()
	if x != t.X || y != t.Y {
		t.Errorf("Coordinates do not match %v %s", t.Point, t.String())
	}
}

func (t OnePoint) is_stringable() {
	t.Helper()
	str := t.String()
	if str == "" {
		t.Errorf("%v is not stringable", t.Point)
	}
}
