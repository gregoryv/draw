package design

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestLink_class(t *testing.T) {
	cases := []struct {
		lnk Link
		exp string
	}{
		{Link{}, "arrow"},
		{Link{Class: "special"}, "special"},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := c.lnk.class()
		assert().Equals(got, c.exp)
	}
}
