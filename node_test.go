package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_rect(t *testing.T) {
	cases := []struct {
		node *Node
		exp  string
	}{
		{NewNode(Element_rect), "<rect>"},
		{
			node: NewNode(
				Element_rect,
				Attribute{"x", "100"},
				Attribute{"y", "0"},
			),
			exp: `<rect x="100" y="0">`,
		},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		buf := bytes.NewBufferString("")
		c.node.WriteTo(buf)
		got := buf.String()
		assert().Equals(got, c.exp)
	}
}
