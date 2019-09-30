package xy

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		l1, l2 *Line
		exp    Position // expected intersection
	}{
		{
			l1:  NewLine(0, 0, 5, 5),
			l2:  NewLine(2, 0, 0, 2),
			exp: Position{1, 1},
		},
		{
			l1:  NewLine(0, 10, 10, 10),
			l2:  NewLine(5, 0, 5, 20),
			exp: Position{5, 10},
		},
	}
	for i, c := range cases {
		testcase := fmt.Sprintf("testcase %v", i)
		t.Run(testcase, func(t *testing.T) {
			t.Logf("Between lines\nl1: %v\nl2: %v", c.l1, c.l2)
			got, err := c.l1.Intersect(c.l2)
			if err != nil {
				t.Error(err)
			}
			if !got.Equals(c.exp) {
				t.Errorf("%v is not %v", got, c.exp)
			}
		})
	}
}
