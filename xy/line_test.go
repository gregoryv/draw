package xy

import (
	"fmt"
	"testing"
)

func TestIntersectSegment(t *testing.T) {
	cases := []struct {
		l1, l2 *Line
		exp    Position // expected intersection
		ok     bool
	}{
		{
			l1:  NewLine(0, 0, 5, 5),
			l2:  NewLine(2, 0, 0, 2),
			exp: Position{1, 1},
			ok:  true,
		},
		{
			l1:  NewLine(0, 10, 10, 10),
			l2:  NewLine(5, 0, 5, 20),
			exp: Position{5, 10},
			ok:  true,
		},
		{
			l1:  NewLine(0, 10, 0, 10),
			l2:  NewLine(5, 10, 5, 10),
			exp: Position{0, 0},
			ok:  false,
		},
		{
			l1:  NewLine(0, 0, 1, 5),
			l2:  NewLine(5, 0, 4, 5),
			exp: Position{0, 0},
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
