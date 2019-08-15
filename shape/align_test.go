package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAlignHorizontal(t *testing.T) {
	cases := []struct {
		shapeA, shapeB Shape // B is aligned to A
		expX, expY     int
	}{
		{
			&Label{X: 10, Y: 10},
			&Label{X: 50, Y: 40},
			50, 10,
		},
		{
			&Label{X: 10, Y: 10},
			&Record{X: 50, Y: 40},
			50, 10,
		},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		AlignHorizontal(Top, c.shapeA, c.shapeB)
		x, y := c.shapeB.Position()
		assert().Equals(x, c.expX)
		assert().Equals(y, c.expY)
	}

}
