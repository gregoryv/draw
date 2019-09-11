package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_one_line(t *testing.T) {
	it := &one_line{t, NewLine(1, 1, 7, 7)}
	it.can_be_rendered_as_svg()
	it.s_directed_right()
	// when
	it = &one_line{t, NewLine(8, 8, 0, 0)}
	it.s_directed_left()
	it.can_move()
	// when
	line := NewLine(0, 0, 0, 0)
	line.Class = "special"
	it = &one_line{t, line}
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

func (t *one_line) s_directed_right() {
	t.Helper()
	dir := t.Direction()
	if dir != LR {
		t.Error("not directed right")
	}
}

func (t *one_line) s_directed_left() {
	t.Helper()
	dir := t.Direction()
	if dir != RL {
		t.Error("not directed left")
	}
}

func (t *one_line) can_move() {
	t.Helper()
	s, e := t.Start, t.End
	t.SetX(s.X + 1)
	t.SetY(s.Y + 1)
	assert := asserter.New(t)
	assert(!t.Start.Equals(s)).Errorf("start position same")
	assert(!t.End.Equals(e)).Errorf("end position same")
}
