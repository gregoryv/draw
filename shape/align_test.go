package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAlignVertical(t *testing.T) {
	var aligner Aligner
	cases := []struct {
		align          func(...Shape)
		shapeA, shapeB Shape // B is aligned to A
		expX, expY     int
	}{
		{
			aligner.VAlignCenter,
			&Label{X: 10, Y: 10, Text: "1234", Font: Font{Width: 10}},
			&Label{X: 50, Y: 40, Text: "12", Font: Font{Width: 10}},
			20, 40,
		},
		{
			aligner.VAlignLeft,
			&Label{X: 10, Y: 10, Text: "1234", Font: Font{Width: 10}},
			&Label{X: 50, Y: 40, Text: "12", Font: Font{Width: 10}},
			10, 40,
		},
		{
			aligner.VAlignRight,
			&Label{X: 10, Y: 10, Text: "1234", Font: Font{Width: 10}},
			&Label{X: 50, Y: 40, Text: "12", Font: Font{Width: 10}},
			30, 40,
		},
		{
			aligner.VAlignCenter,
			NewLine(10, 10, 4, 10), // Right to left
			NewLine(0, 0, 10, 0),
			2, 0,
		},
	}
	assert := asserter.New(t)
	for i, c := range cases {
		c.align(c.shapeA, c.shapeB)
		x, y := c.shapeB.Position()
		assert(x == c.expX).Errorf("%v. X was %v, expected %v", i, x, c.expX)
		assert(y == c.expY).Errorf("%v. Y was %v, expected %v", i, y, c.expY)
	}
}

func TestAlignHorizontal(t *testing.T) {
	var aligner Aligner
	cases := []struct {
		align          func(...Shape)
		shapeA, shapeB Shape // B is aligned to A
		expX, expY     int
	}{
		{
			aligner.HAlignTop,
			&Label{X: 10, Y: 10},
			&Label{X: 50, Y: 40},
			50, 10,
		},
		{
			aligner.HAlignBottom,
			&Label{X: 10, Y: 10},
			&Record{X: 50, Y: 40},
			50, 10,
		},
		{
			aligner.HAlignCenter,
			NewLine(0, 10, 0, 20),
			NewLine(0, 20, 0, 30),
			0, 10,
		},
		{
			aligner.HAlignCenter,
			NewLine(0, 20, 0, 30),
			NewLine(0, 10, 0, 20),
			0, 20,
		},
		{ // first line is shorter
			aligner.HAlignCenter,
			NewLine(0, 10, 0, 15),
			NewLine(0, 10, 0, 20),
			0, 8,
		},
		{ // second line is shorter
			aligner.HAlignCenter,
			NewLine(0, 10, 0, 20),
			NewLine(0, 10, 0, 15),
			0, 12,
		},
		{
			aligner.HAlignCenter,
			NewLine(0, 10, 0, 20),
			&Label{
				X: 0, Y: 20, Font: Font{Height: 10},
			},
			0, 20,
		},
		{
			aligner.HAlignCenter,
			&Label{X: 0, Y: 20, Font: Font{Height: 10}},
			&Label{X: 0, Y: 20, Font: Font{Height: 6}},
			0, 28,
		},
	}
	assert := asserter.New(t)
	for i, c := range cases {
		c.align(c.shapeA, c.shapeB)
		x, y := c.shapeB.Position()
		assert(x == c.expX).Errorf("%v. X was %v, expected %v", i, x, c.expX)
		assert(y == c.expY).Errorf("%v. Y was %v, expected %v", i, y, c.expY)
	}
}
