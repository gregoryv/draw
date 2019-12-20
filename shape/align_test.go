package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw/xy"
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
				pos:  xy.Position{10, 10},
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				pos:  xy.Position{50, 40},
				Text: "12",
				Font: DefaultFont,
			},
			17, 40,
		},
		{
			aligner.VAlignLeft,
			&Label{
				pos:  xy.Position{10, 10},
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				pos:  xy.Position{50, 40},
				Text: "12",
				Font: DefaultFont,
			},
			10, 40,
		},
		{
			aligner.VAlignRight,
			&Label{
				pos:  xy.Position{10, 10},
				Text: "1234",
				Font: DefaultFont,
			},
			&Label{
				pos:  xy.Position{50, 40},
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
			&Label{pos: xy.Position{10, 10}},
			&Label{pos: xy.Position{50, 40}},
			50, 10,
		},
		{
			aligner.HAlignBottom,
			&Label{pos: xy.Position{10, 10}},
			&Label{pos: xy.Position{50, 40}},
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
				pos:  xy.Position{0, 20},
				Font: Font{LineHeight: 10},
			},
			0, 20,
		},
		{
			aligner.HAlignCenter,
			&Label{
				pos:  xy.Position{0, 20},
				Font: Font{LineHeight: 10},
			},
			&Label{
				pos:  xy.Position{0, 20},
				Font: Font{LineHeight: 6},
			},
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
