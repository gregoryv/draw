package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestOneLabel(t *testing.T) {
	it := NewOneLabel(t)
	// when
	it.IsEmpty()
	it.RendersAsSvg()
	it.HasNoWidth()
	// when
	it.IsNotEmpty()
	it.IsStyled()
	it.HasWidth()
	it.CanMove()
}

func NewOneLabel(t *testing.T) *OneLabel {
	return &OneLabel{
		T:      t,
		assert: asserter.New(t),
	}
}

type OneLabel struct {
	*testing.T
	assert
	*Label
}

type assert = asserter.AssertFunc

func (t *OneLabel) IsEmpty() {
	t.Label = NewLabel("")
}

func (t *OneLabel) IsNotEmpty() {
	t.Label = &Label{Text: "a label"}
}

func (t *OneLabel) RendersAsSvg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	t.assert().Contains(buf.String(), "<text ")
}

func (t *OneLabel) HasNoWidth() {
	t.Helper()
	t.assert(t.Width() == 0).Error("has width")
}

func (t *OneLabel) IsStyled() {
	t.Font = Font{Height: 9, Width: 7, LineHeight: 15}
	t.Pad = Padding{3, 3, 10, 2}
}

func (t *OneLabel) HasWidth() {
	t.Helper()
	t.assert(t.Width() > 0).Error("0 width")
}

func (t *OneLabel) CanMove() {
	t.Helper()
	x, y := t.Position()
	t.SetX(x + 1)
	t.SetY(y + 1)
	dir := t.Direction()
	t.assert(dir == LR).Error("Direction should always be LR for record")
	t.assert(x != t.X).Error("x still the same")
	t.assert(y != t.Y).Error("y still the same")
}
