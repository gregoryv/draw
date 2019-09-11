package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_one_line(t *testing.T) {
	it := &one_line{t, NewLine(1, 1, 7, 7)}
	it.can_be_rendered_as_svg()
}

type one_line struct {
	*testing.T
	*Line
}

func (t *one_line) can_be_rendered_as_svg() {
	t.Helper()
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	assert := asserter.New(t)
	assert().Contains(buf.String(), "<line ")
}
