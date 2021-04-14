package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestOneRecord(t *testing.T) {
	it := NewOneRecord(t)
	it.SHeightAdapts()
	it.SWidthAdapts()
	it.RendersAsSvg()
	it.CanHideFields()
	it.CanHideMethods()

	it = NewOneRecord(t)
	it.HideFields()
	it.IsMissingFields()
}

func NewOneRecord(t *testing.T) *OneRecord {
	rec := NewRecord("car")
	rec.Fields = []string{"short", "longerField"}
	rec.Methods = []string{"String", "Model"}
	return &OneRecord{t, rec, asserter.New(t)}
}

type OneRecord struct {
	*testing.T
	*Record
	assert
}

func (t *OneRecord) CanHideFields() {
	t.HideFields()
	t.assert(len(t.Fields) == 0).Error("fields not hidden")
}

func (t *OneRecord) IsMissingFields() {
	t.assert(len(t.Fields) == 0).Error("has fields")
}

func (t *OneRecord) CanHideMethods() {
	t.assert(!t.HideMethod("no-such-method")).Error("found non existing method")
	m := t.Methods[0]
	t.assert(t.HideMethod(m)).Errorf("method %q not hidden", m)
	t.HideMethods()
	t.assert(len(t.Methods) == 0).Error("methods not hidden")
}

func (t *OneRecord) SHeightAdapts() {
	t.assert(t.Height() > 0).Error("height did not adapt")
}

func (t *OneRecord) SWidthAdapts() {
	t.assert(t.Width() > 0).Error("width did not adapt")
}

func (t *OneRecord) RendersAsSvg() {
	buf := &bytes.Buffer{}
	t.WriteSVG(buf)
	t.assert().Contains(buf, "<rect ")
}
