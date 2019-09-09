package style

import (
	"bytes"
	"testing"
)

func Test_one_styler(t *testing.T) {
	it := one_styler{T: t}

	it.does_not_modify_unrecognized_elements()
	it.skips_non_classed_elements()
	it.only_modifies_classed_elements()
	it.rejects_bad_elements()
}

type one_styler struct {
	*testing.T
}

func (t *one_styler) rejects_bad_elements() {
	t.Helper()
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Styler should panic on malformed xml")
		}
	}()
	buf := &bytes.Buffer{}
	s := NewStyler(buf)
	input := `<line class=">`
	s.Write([]byte(input))
	if buf.String() != input {
		t.Error(buf.String())
	}
}

func (t *one_styler) skips_non_classed_elements() {
	t.Helper()
	buf := &bytes.Buffer{}
	s := NewStyler(buf)
	input := `<line />`
	s.Write([]byte(input))
	if buf.String() != input {
		t.Fail()
	}
}

func (t *one_styler) does_not_modify_unrecognized_elements() {
	t.Helper()
	buf := &bytes.Buffer{}
	s := NewStyler(buf)
	input := `<x class="not-something-we-recognize" />"`
	s.Write([]byte(input))
	if buf.String() != input {
		t.Fail()
	}
}

func (t *one_styler) only_modifies_classed_elements() {
	t.Helper()
	buf := &bytes.Buffer{}
	s := NewStyler(buf)
	input := `<x class="line" />"`
	s.Write([]byte(input))
	if buf.String() == input {
		t.Fail()
	}
}
