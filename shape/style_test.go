package shape

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestStyler_rejects_bad_elements(t *testing.T) {
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Styler should panic on malformed xml")
		}
	}()
	buf := &bytes.Buffer{}
	s := NewStyle(buf)
	input := `<line class=">`
	s.Write([]byte(input))
	if buf.String() != input {
		t.Error(buf.String())
	}
}

func TestStylerWrite_adds_style_to_classed_elements(t *testing.T) {
	cases := []struct {
		input string
		exp   string
	}{
		{
			`<x class="line" />`,
			`<x stroke="#d3d3d3" />`,
		},
		{
			`<x class="whatever" />`,
			`<x class="whatever" />`,
		},
		{
			`<x b="t" s="x"/>`,
			`<x b="t" s="x"/>`,
		},
		{
			`<x />`,
			`<x />`,
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			var buf bytes.Buffer
			s := NewStyle(&buf)
			s.Write([]byte(c.input))
			got := buf.String()
			assert := asserter.New(t)
			assert().Equals(got, c.exp)
		})
	}
}

func TestStyler_write(t *testing.T) {
	s := &Style{
		err: fmt.Errorf("wrong"),
	}
	s.write([]byte("something"))
	if s.written > 0 {
		t.Errorf("Wrote %v, should be 0", s.written)
	}

}
