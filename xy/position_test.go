package xy

import (
	"testing"
)

func TestPosition_Distance(t *testing.T) {
	var p Position
	cases := []struct {
		q   Position
		exp float64
	}{
		{Position{0, 2}, 2},
	}
	for _, c := range cases {
		got := p.Distance(c.q)
		if got != c.exp {
			t.Error(got, c.exp)
		}
	}

}

func Test_one_position(t *testing.T) {
	it := one_position{t, Position{10, 10}}
	q := Position{10, 10}
	it.is_not_left_of(q)
	it.is_not_right_of(q)
	it.is_not_above(q)
	it.is_not_below(q)
	it.is_same_as(q)
	it.has_quick_access_to_coordinates()
	it.is_stringable()
}

type one_position struct {
	*testing.T
	Position
}

func (t one_position) is_left_of(q Position) {
	t.Helper()
	if !t.LeftOf(q) {
		t.Errorf("%v should be left of %v", t.Position, q)
	}
}

func (t one_position) is_not_left_of(q Position) {
	t.Helper()
	if t.LeftOf(q) {
		t.Errorf("%v should not be left of %v", t.Position, q)
	}
}

func (t one_position) is_not_right_of(q Position) {
	t.Helper()
	if t.RightOf(q) {
		t.Errorf("%v should not be right of %v", t.Position, q)
	}
}

func (t one_position) is_not_above(q Position) {
	t.Helper()
	if t.Above(q) {
		t.Errorf("%v should not be above %v", t.Position, q)
	}
}

func (t one_position) is_not_below(q Position) {
	t.Helper()
	if t.Below(q) {
		t.Errorf("%v should not be below %v", t.Position, q)
	}
}

func (t one_position) is_same_as(q Position) {
	t.Helper()
	if !t.Equals(q) {
		t.Errorf("%v should be equal to %v", t.Position, q)
	}
}

func (t one_position) has_quick_access_to_coordinates() {
	t.Helper()
	x, y := t.XY()
	if x != t.X || y != t.Y {
		t.Errorf("Coordinates do not match %v %s", t.Position, t.String())
	}
}

func (t one_position) is_stringable() {
	t.Helper()
	str := t.String()
	if str == "" {
		t.Errorf("%v is not stringable", t.Position)
	}
}
