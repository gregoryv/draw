package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_one_label(t *testing.T) {
	it := new_one_label(t)
	// when
	it.is_empty()
	it.can_be_rendered_as_svg()
	it.has_no_width()
	// when
	it.is_not_empty()
	it.is_styled()
	it.has_width()
	it.can_move()
}

func new_one_label(t *testing.T) *one_label {
	return &one_label{
		T:      t,
		assert: asserter.New(t),
	}
}

type one_label struct {
	*testing.T
	assert
	*Label
}

type assert = asserter.AssertFunc

func (t *one_label) is_empty() {
	t.Label = &Label{}
}

func (t *one_label) is_not_empty() {
	t.Label = &Label{Text: "a label"}
}

func (t *one_label) can_be_rendered_as_svg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	t.assert().Contains(buf.String(), "<text ")
}

func (t *one_label) has_no_width() {
	t.Helper()
	t.assert(t.Width() == 0).Error("has width")
}

func (t *one_label) is_styled() {
	t.Font = Font{Height: 9, Width: 7, LineHeight: 15}
	t.Pad = Padding{3, 3, 10, 2}
}

func (t *one_label) has_width() {
	t.Helper()
	t.assert(t.Width() > 0).Error("0 width")
}

func (t *one_label) can_move() {
	t.Helper()
	x, y := t.Position()
	t.SetX(x + 1)
	t.SetY(y + 1)
	dir := t.Direction()
	t.assert(dir == LR).Error("Direction should always be LR for record")
	t.assert(x != t.X).Error("x still the same")
	t.assert(y != t.Y).Error("y still the same")
}
