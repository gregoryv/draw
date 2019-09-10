package shape

import (
	"os"
	"testing"

	"github.com/gregoryv/go-design/style"
)

func Test_one_record(t *testing.T) {
	it := &one_record{
		T: t,
	}
	it.reflects_a_struct()
	it.is_styled()
	it.has_title_and_fields()
}

type one_record struct {
	*testing.T
	*Record
}

func (t *one_record) reflects_a_struct() {
	t.Record = NewStructRecord(Record{})
}

func (t *one_record) is_styled() {
	t.Font = Font{Height: 9, Width: 7, LineHeight: 15}
}

func (t *one_record) has_title_and_fields() {
	t.saveAs("img/record_with_title_and_fields.svg")
}

func (t *one_record) saveAs(filename string) {
	t.Helper()
	d := &Svg{Width: 200, Height: 300}
	d.Append(t.Record)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(style.NewStyler(fh))
	fh.Close()
}
