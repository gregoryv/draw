package design

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/golden"
)

func TestSequenceDiagram_AddInterface(t *testing.T) {
	d := NewSequenceDiagram()
	type X interface{}
	got := d.AddInterface((*X)(nil))
	exp := "design.X"
	assert := asserter.New(t)
	assert().Equals(got, exp)

}

func TestSequenceDiagram_WithCaption(t *testing.T) {
	d := NewSequenceDiagram()
	before := d.Width()
	d.SetCaption("should affect width")
	d.WriteSVG(ioutil.Discard) // caption is not added until written
	after := d.Width()
	if before == after {
		t.Fail()
	}
}

func TestSequenceDiagram_Inline(t *testing.T) {
	var (
		d = NewSequenceDiagram()
		a = d.Add("a")
		b = d.Add("b")
	)
	d.Link(a, b, "text")
	d.SetCaption("should affect width")
	got := d.Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
}

func TestSequenceDiagram_String(t *testing.T) {
	var (
		d = NewSequenceDiagram()
		a = d.Add("a")
		b = d.Add("b")
	)
	d.Link(a, b, "text")
	d.SetCaption("should affect width")
	got := d.String()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
}

func TestSequenceDiagram_AddStruct(t *testing.T) {
	d := NewSequenceDiagram()
	before := d.Width()
	d.AddStruct(d)
	after := d.Width()
	if before == after {
		t.Fail()
	}
}

func Test_one_sequence_diagram(t *testing.T) {
	it := new_one_sequence_diagram(t)
	it.is_empty()
	// when
	it.has_columns()
	it.has_no_links()
	it.is_not_empty()
	// and if
	it.has_fixed_size()
	it.is_rendered_with_fixed_size()
}

func new_one_sequence_diagram(t *testing.T) *one_sequence_diagram {
	return &one_sequence_diagram{t, NewSequenceDiagram()}
}

type one_sequence_diagram struct {
	*testing.T
	*SequenceDiagram
}

func (t *one_sequence_diagram) has_columns()    { t.AddColumns("a", "b") }
func (t *one_sequence_diagram) has_no_links()   { t.ClearLinks() }
func (t *one_sequence_diagram) has_fixed_size() { t.SetWidth(10); t.SetHeight(10) }

func (t *one_sequence_diagram) is_empty() {
	t.Helper()
	w := bytes.NewBufferString("")
	t.WriteSVG(w)
	golden.Assert(t, w.String())
}

func (t *one_sequence_diagram) is_not_empty() {
	t.Helper()
	assert := asserter.New(t)
	assert(t.Width() != 0).Error("Width is 0")
	assert(t.Height() != 0).Error("Height is 0")
}

func (t *one_sequence_diagram) is_rendered_with_fixed_size() {
	t.Helper()
	w := bytes.NewBufferString("")
	t.WriteSVG(w)
	golden.Assert(t, w.String())
}
