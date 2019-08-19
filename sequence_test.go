package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestSequenceDiagram_Height(t *testing.T) {
	fixed := SequenceDiagram{}
	fixed.SetHeight(100)
	fixed.SetWidth(200)
	cases := []struct {
		dia                 SequenceDiagram
		expHeight, expWidth int
	}{
		{SequenceDiagram{}, 0, 0},
		{fixed, 100, 200},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		assert().Equals(c.dia.Height(), c.expHeight)
		assert().Equals(c.dia.Width(), c.expWidth)
	}

}
