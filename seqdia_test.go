package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/golden"
)

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
	t.WriteSvg(w)
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
	t.WriteSvg(w)
	golden.Assert(t, w.String())
}
