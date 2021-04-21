package shape

import (
	"fmt"
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
			&Label{
				x:    10,
				y:    10,
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				x:    50,
				y:    40,
				Text: "12",
				Font: DefaultFont,
			},
			17, 40,
		},
		{
			aligner.VAlignLeft,
			&Label{
				x:    10,
				y:    10,
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				x:    50,
				y:    40,
				Text: "12",
				Font: DefaultFont,
			},
			10, 40,
		},
		{
			aligner.VAlignRight,
			&Label{
				x:    10,
				y:    10,
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				x:    50,
				y:    40,
				Text: "12",
				Font: DefaultFont,
			},
			24, 40,
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
			&Label{Font: DefaultFont, x: 10, y: 10},
			&Label{Font: DefaultFont, x: 50, y: 40},
			50, 10,
		},
		{
			aligner.HAlignBottom,
			&Label{x: 10, y: 10},
			&Label{x: 50, y: 40},
			50, 10,
		},
		{
			aligner.HAlignCenter,
			&Label{Font: DefaultFont, x: 10, y: 10},
			&Label{Font: DefaultFont, x: 50, y: 40},
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
				y:    20,
				Font: Font{LineHeight: 10},
			},
			0, 10,
		},
		{
			aligner.HAlignCenter,
			&Label{
				y:    20,
				Font: Font{LineHeight: 10},
			},
			&Label{
				y:    20,
				Font: Font{LineHeight: 6},
			},
			0, 22,
		},
	}

	for i, c := range cases {
		name := fmt.Sprintf("%v %v", c.shapeA, c.shapeB)
		t.Run(name, func(t *testing.T) {
			assert := asserter.New(t)

			c.align(c.shapeA, c.shapeB)
			x, y := c.shapeB.Position()
			assert(x == c.expX).Errorf("%v. X was %v, expected %v", i, x, c.expX)
			assert(y == c.expY).Errorf("%v. Y was %v, expected %v", i, y, c.expY)
		})
	}
}
