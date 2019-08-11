package svg

import (
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/go-design/xml"
)

func Test_element_names(t *testing.T) {
	for e := Element_undefined; e <= Element_last+1; e++ {
		valid := e > Element_undefined && e < Element_last
		name := e.String()
		if valid && name == "undefined" {
			t.Errorf("No names for %#v", e)
		}
	}
}

func Test_xml_elements(t *testing.T) {
	cases := []struct {
		n   *xml.Node
		exp string
	}{
		{Rect(0, 0, 1, 1), "<rect"},
		{Line(0, 0, 1, 1), "<line x1"},
		{Text(0, 0, "hello"), "<text"},
	}
	assert := asserter.New(t)

	for _, c := range cases {
		got := c.n.String()
		assert().Contains(got, c.exp)
	}

}
