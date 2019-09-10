package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_one_record(t *testing.T) {
	rec := NewRecord("car")
	rec.Fields = []string{"short", "longerField"}
	rec.Methods = []string{"String", "Model"}
	it := &one_record{t, rec}
	it.has_fields()
	it.has_methods()
	it.is_styled()
	it.s_height_adapts()
	it.s_width_adapts()
	it.can_be_rendered_as_svg()
	it.can_move()

	it = &one_record{t, NewStructRecord(Record{})}
	it.has_fields()
	it.has_methods()

	it = &one_record{t, NewInterfaceRecord((*Shape)(nil))}
	it.is_missing_fields()
	it.has_methods()

	it = &one_record{t, NewRecord("simple")}
	it.is_missing_fields()
	it.is_styled()
	it.s_height_adapts()
	it.s_width_adapts()
}

type one_record struct {
	*testing.T
	*Record
}

func (t *one_record) can_move() {
	t.Helper()
	t.SetX(10)
	dir := t.Direction()
	assert := asserter.New(t)
	assert(dir == LR).Error("Direction should always be LR for record")
}

func (t *one_record) has_fields() {
	t.Helper()
	assert := asserter.New(t)
	assert(len(t.Fields) >= 0).Error("missing fields")
}

func (t *one_record) is_missing_fields() {
	t.Helper()
	assert := asserter.New(t)
	assert(len(t.Fields) == 0).Error("has fields")
}

func (t *one_record) has_methods() {
	t.Helper()
	assert := asserter.New(t)
	assert(len(t.Methods) > 0).Error("missing methods")
}

func (t *one_record) s_height_adapts() {
	t.Helper()
	assert := asserter.New(t)
	assert(t.Height() > 0).Error("height did not adapt")
}

func (t *one_record) s_width_adapts() {
	t.Helper()
	assert := asserter.New(t)
	assert(t.Width() > 0).Error("width did not adapt")
}

func (t *one_record) is_styled() {
	t.SetFont(Font{Height: 9, Width: 7, LineHeight: 15})
	t.SetTextPad(Padding{3, 3, 10, 2})
}

func (t *one_record) reflects_an_interface() {
	t.Record = NewInterfaceRecord((*Shape)(nil))
}

func (t *one_record) can_be_rendered_as_svg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	assert := asserter.New(t)
	assert().Contains(buf.String(), "<rect ")
}
