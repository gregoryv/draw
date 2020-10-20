package draw

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestStyle_CSS(t *testing.T) {
	css := DefaultClassAttributes.CSS()
	if !strings.Contains(css, ".database") {
		t.Error("missing .database\n", css)
	}
	if !strings.Contains(css, `stroke="black";\n`) {
		t.Error(`missing stroke="black";`, css)
	}
}

func TestStyle_rejects_bad_elements(t *testing.T) {
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

func TestStyle_Write_adds_style_to_classed_elements(t *testing.T) {
	cases := []struct {
		input string
		exp   string
	}{
		{
			`<x class="line" />`,
			`<x stroke="black" />`,
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

func TestStyle_write(t *testing.T) {
	s := &Style{
		err: fmt.Errorf("wrong"),
	}
	s.write([]byte("something"))
	if s.written > 0 {
		t.Errorf("Wrote %v, should be 0", s.written)
	}
}

func TestStyle_SetOutput(t *testing.T) {
	s := NewStyle(nil)
	s.SetOutput(nil)
}
