package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestOneRecord(t *testing.T) {
	rec := NewRecord("car")
	rec.Fields = []string{"short", "longerField"}
	rec.Methods = []string{"String", "Model"}
	it := NewOneRecord(t, rec)
	it.HasFields()
	it.HasMethods()
	it.IsStyled()
	it.SHeightAdapts()
	it.SWidthAdapts()
	it.RendersAsSvg()
	it.CanMove()

	it = NewOneRecord(t, NewStructRecord(Record{}))
	it.HasFields()
	it.HasMethods()
	it.CanHideFields()
	it.CanHideMethods()

	it = NewOneRecord(t, NewInterfaceRecord((*Shape)(nil)))
	it.IsMissingFields()
	it.HasMethods()

	it = NewOneRecord(t, NewRecord("simple"))
	it.IsMissingFields()
	it.IsStyled()
	it.SHeightAdapts()
	it.SWidthAdapts()
}

func NewOneRecord(t *testing.T, rec *Record) *OneRecord {
	return &OneRecord{t, asserter.New(t), rec}
}

type OneRecord struct {
	*testing.T
	assert
	*Record
}

func (t *OneRecord) CanMove() {
	t.Helper()
	t.SetX(10)
	dir := t.Direction()
	t.assert(dir == LR).Error("Direction should always be LR for record")
}

func (t *OneRecord) HasFields() {
	t.Helper()
	t.assert(len(t.Fields) >= 0).Error("missing fields")
}

func (t *OneRecord) CanHideFields() {
	t.Helper()
	t.HideFields()
	t.assert(len(t.Fields) == 0).Error("fields not hidden")
}

func (t *OneRecord) IsMissingFields() {
	t.Helper()
	t.assert(len(t.Fields) == 0).Error("has fields")
}

func (t *OneRecord) HasMethods() {
	t.Helper()
	t.assert(len(t.Methods) > 0).Error("missing methods")
}

func (t *OneRecord) CanHideMethods() {
	t.Helper()
	t.HideMethods()
	t.assert(len(t.Methods) == 0).Error("methods not hidden")
}

func (t *OneRecord) SHeightAdapts() {
	t.Helper()
	t.assert(t.Height() > 0).Error("height did not adapt")
}

func (t *OneRecord) SWidthAdapts() {
	t.Helper()
	t.assert(t.Width() > 0).Error("width did not adapt")
}

func (t *OneRecord) IsStyled() {
	t.SetFont(Font{Height: 9, Width: 7, LineHeight: 15})
	t.SetTextPad(Padding{3, 3, 10, 2})
}

func (t *OneRecord) RendersAsSvg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	assert := asserter.New(t)
	assert().Contains(buf.String(), "<rect ")
}
