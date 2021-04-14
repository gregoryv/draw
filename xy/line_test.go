package xy

import (
	"fmt"
	"testing"
)

func TestIntersectSegment(t *testing.T) {
	hLine := NewLine(5, 5, 15, 5)
	vLine := NewLine(10, 0, 10, 10)
	cases := []struct {
		l1, l2 *Line
		exp    Point // expected intersection
		ok     bool
	}{
		{
			// Crossed as plus
			l1:  hLine,
			l2:  vLine,
			exp: Point{10, 5},
			ok:  true,
		},
		{
			// from left
			l1:  NewLine(0, 0, 11, 11),
			l2:  vLine,
			exp: Point{10, 10},
			ok:  true,
		},
		{
			// from right
			l1:  NewLine(15, 0, 0, 10),
			l2:  vLine,
			exp: Point{10, 3},
			ok:  true,
		},
		{
			// from right
			l1:  NewLine(15, 0, 13, 0),
			l2:  vLine,
			exp: Point{10, 0},
			ok:  false,
		},
		{
			// from above
			l1:  NewLine(0, 0, 12, 10),
			l2:  hLine,
			exp: Point{6, 5},
			ok:  true,
		},
		{
			// from below
			l1:  NewLine(12, 10, 0, 0),
			l2:  hLine,
			exp: Point{6, 5},
			ok:  true,
		},
		{
			l1:  NewLine(0, 10, 0, 0), // vertical upward
			l2:  NewLine(11, 0, -1, 11),
			exp: Point{0, 10},
			ok:  true,
		},
		{
			l1:  NewLine(0, 10, 0, 10), // point
			l2:  NewLine(5, 10, 5, 10), // point
			exp: Point{0, 0},
			ok:  false,
		},
	}
	for i, c := range cases {
		testcase := fmt.Sprintf("testcase %v", i)
		t.Run(testcase, func(t *testing.T) {
			t.Logf("Between lines\nl1: %v\nl2: %v", c.l1, c.l2)
			got, err := c.l1.IntersectSegment(c.l2)
			if c.ok && err != nil {
				t.Error(err)
			}
			if !c.ok && err == nil {
				t.Error("Ought to fail")
			}

			if !got.Equals(c.exp) {
				t.Errorf("%v is not %v", got, c.exp)
			}
		})
	}
}
