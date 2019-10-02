package shape

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestLine(t *testing.T) {
	testShape(t, NewLine(1, 1, 7, 7))
}

func TestOneLine(t *testing.T) {
	it := &OneLine{t, NewLine(1, 1, 7, 7)}
	it.RendersAsSvg()
	it.HasDirection()
	// when
	it = &OneLine{t, NewLine(8, 8, 0, 0)}
	it.CanMove()
	it.HasWidth()
	// when
	line := NewLine(0, 0, 0, 0)
	line.Class = "special"
	it = &OneLine{t, line}
	it.RendersAsSvg()
}

type OneLine struct {
	*testing.T
	*Line
}

func (t *OneLine) RendersAsSvg() {
	t.Log("Renders as SVG")
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	assert := asserter.New(t)
	assert().Contains(buf.String(), "<line ")
}

func (t *OneLine) HasDirection() {
	t.Log("Has direction")
	line := NewLine(1, 1, 7, 7)
	dir := line.Direction()
	assert := asserter.New(t)
	assert(dir == LR).Errorf("%v shold have LR was %v", line, dir)

	line = NewLine(8, 8, 0, 0)
	dir = line.Direction()
	assert(dir == RL).Errorf("%v shold have RL was %v", line, dir)
}

func (t *OneLine) CanMove() {
	t.Helper()
	s, e := t.Start, t.End
	t.SetX(s.X + 1)
	t.SetY(s.Y + 1)
	assert := asserter.New(t)
	assert(!t.Start.Equals(s)).Errorf("start position same")
	assert(!t.End.Equals(e)).Errorf("end position same")
}

func (t *OneLine) HasWidth() {
	t.Helper()
	assert := asserter.New(t)
	assert(t.Width() > 0).Error("0 width")
}
