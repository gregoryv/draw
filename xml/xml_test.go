package xml

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_attribute(t *testing.T) {
	cases := []struct {
		attr Attribute
		exp  string
	}{
		{NewAttribute("src", "http"), `src="http"`},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		got := c.attr.String()
		assert().Equals(got, c.exp)
	}
}

func Test_drawables(t *testing.T) {
	cases := []struct {
		node Drawable
		exp  string
	}{
		{CData("hello"), "hello"},
		{NewNode(&img{}), "<img/>\n"},
		{
			node: NewNode(
				&img{},
				Attribute{"src", "http://example.com"},
			),
			exp: "<img src=\"http://example.com\"/>\n",
		},
		{
			node: &Node{
				element: &img{},
				children: []Drawable{
					NewNode(
						&img{},
						Attribute{"src", "http://example.com"},
					),
				},
			},
			exp: "<img><img src=\"http://example.com\"/>\n</img>",
		},
	}
	assert := asserter.New(t)
	for _, c := range cases {
		buf := bytes.NewBufferString("")
		c.node.WriteTo(buf)
		got := buf.String()
		assert().Equals(got, c.exp)
		n, ok := c.node.(Stringer)
		if ok {
			assert().Equals(n.String(), c.exp)
		}
	}
}

type Stringer interface {
	String() string
}

type img struct{}

func (*img) String() string {
	return "img"
}
