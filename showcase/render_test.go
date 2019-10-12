package showcase

import (
	"fmt"
	"testing"
)

func TestRenderAllDiagrams(t *testing.T) {
	cases := []struct {
		d        saveCap
		filename string
		caption  string
	}{
		{BasicNetHttpClassDiagram(), "nethttp.svg", "ServeMux routes requests to handlers"},
		{BackendHandler(), "nethttp_backend.svg", "HTTP backends implement handlers"},
	}
	for i, c := range cases {
		c.d.SetCaption(fmt.Sprintf("Figure %v. %s", i+1, c.caption))
		c.d.SaveAs(c.filename)
	}
}

type saveCap interface {
	SetCaption(string)
	SaveAs(string) error
}
