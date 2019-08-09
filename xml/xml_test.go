package xml

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
		{NewNode(&img{}), "<img>"},
		{
			node: NewNode(
				&img{},
				Attribute{"src", "http://example.com"},
			),
			exp: `<img src="http://example.com">`,
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

type img struct{}

func (*img) String() string {
	return "img"
}
