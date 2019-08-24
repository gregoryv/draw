package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAlignHorizontal(t *testing.T) {
	cases := []struct {
		adjust         Adjust
		shapeA, shapeB Shape // B is aligned to A
		expX, expY     int
	}{
		{
			Top,
			&Label{X: 10, Y: 10},
			&Label{X: 50, Y: 40},
			50, 10,
		},
		{
			Top,
			&Label{X: 10, Y: 10},
			&Record{X: 50, Y: 40},
			50, 10,
		},
		{
			Center,
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 20},
			&Line{X1: 0, Y1: 20, X2: 0, Y2: 30},
			0, 10,
		},
		{
			Center,
			&Line{X1: 0, Y1: 20, X2: 0, Y2: 30},
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 20},
			0, 20,
		},
		{ // first line is shorter
			Center,
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 15},
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 20},
			0, 8,
		},
		{ // second line is shorter
			Center,
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 20},
			&Line{X1: 0, Y1: 10, X2: 0, Y2: 15},
			0, 12,
		},
		{
			Center,
			&Line{
				X1: 0, Y1: 10,
				X2: 0, Y2: 20,
			},
			&Label{
				X: 0, Y: 20, Font: Font{Height: 10},
			},
			0, 20,
		},
	}
	assert := asserter.New(t)
	for i, c := range cases {
		HAlign(c.adjust, c.shapeA, c.shapeB)
		x, y := c.shapeB.Position()
		assert(x == c.expX).Errorf("%v. X was %v, expected %v", i, x, c.expX)
		assert(y == c.expY).Errorf("%v. Y was %v, expected %v", i, y, c.expY)
	}

}
