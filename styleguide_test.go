package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_styleguide(t *testing.T) {
	cases := []struct {
		s         StyleGuide
		lines     int
		expHeight int
	}{
		{DefaultStyle, 1, 22},
		{DefaultStyle, 2, 44},
	}
	assert := asserter.New(t)

	for _, c := range cases {
		got := c.s.Height(c.lines)
		assert().Equals(got, c.expHeight)
	}

}
