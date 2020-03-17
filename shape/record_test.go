package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestRecord_HideMethods(t *testing.T) {
	r := NewRecord("a")
	r.Methods = []string{"dddd", "eeeee"}
	r.HideMethods()
	if len(r.Methods) > 0 {
		t.Error("Record still has methods")
	}
}

func TestRecord_HideFields(t *testing.T) {
	r := NewRecord("a")
	r.Fields = []string{"dddd", "eeeee"}
	r.HideFields()
	if len(r.Fields) > 0 {
		t.Error("Record still has fields")
	}
}

func TestRecord_HasMethod(t *testing.T) {
	r := NewRecord("a")
	r.Methods = []string{"b", "c"}
	assert := asserter.New(t)
	assert(r.HideMethod("b")).Error("Existing method b not found")
	assert(!r.HideMethod("x")).Error("Found non existing method c")
}

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
	t.SetFont(Font{Height: 9, LineHeight: 15})
	t.SetTextPad(Padding{3, 3, 10, 2})
}

func (t *OneRecord) RendersAsSvg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	assert := asserter.New(t)
	assert().Contains(buf.String(), "<rect ")
}
