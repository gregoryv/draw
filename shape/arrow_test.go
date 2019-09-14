package shape

import (
	"bytes"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/go-design/style"
)

func Test_one_arrow(t *testing.T) {
	it := new_one_arrow(t)
	it.can_point_up_and_right()
	it.can_point_up_and_left()
	it.can_point_down_and_left()
	it.can_point_down_and_right()
	// also
	it.can_point_right()
	it.can_point_left()
	it.can_point_down()
	it.can_point_up()
	// when
	it.has_a_tail()
	it.has_both_tail_and_head()

	it.can_have_a_specific_class()
	it.can_move()
	it.is_visible()
}

func new_one_arrow(t *testing.T) *one_arrow {
	return &one_arrow{
		T:      t,
		assert: asserter.New(t),
		Arrow:  NewArrow(50, 50, 50, 50),
	}
}

type one_arrow struct {
	*testing.T
	assert
	*Arrow
}

func (t *one_arrow) can_point_up_and_right() {
	t.End.X = t.Start.X + 30
	t.End.Y = t.Start.Y - 30
	t.saveAs("testdata/arrow_points_up_and_right.svg")
}

func (t *one_arrow) can_point_up_and_left() {
	t.End.X = t.Start.X - 30
	t.End.Y = t.Start.Y - 30
	t.saveAs("testdata/arrow_points_up_and_left.svg")
}

func (t *one_arrow) can_point_down_and_left() {
	t.End.X = t.Start.X - 10
	t.End.Y = t.Start.Y + 30
	dir := t.Direction()
	t.assert(dir == RL).Errorf("Direction not right to left: %v", dir)
	t.saveAs("testdata/arrow_points_down_and_left.svg")
}

func (t *one_arrow) can_point_down_and_right() {
	t.End.X = t.Start.X + 20
	t.End.Y = t.Start.Y + 30
	dir := t.Direction()
	t.assert(dir == LR).Errorf("Direction not left to right: %v", dir)
	t.saveAs("testdata/arrow_points_down_and_right.svg")
}

func (t *one_arrow) can_point_right() {
	t.End.X = t.Start.X + 50
	t.End.Y = t.Start.Y
	dir := t.Direction()
	t.assert(dir == LR).Errorf("Direction not left to right: %v", dir)
	t.saveAs("testdata/arrow_points_right.svg")
}

func (t *one_arrow) can_point_left() {
	t.End.X = t.Start.X - 40
	t.End.Y = t.Start.Y
	dir := t.Direction()
	t.assert(dir == RL).Errorf("Direction not right to left: %v", dir)
	t.saveAs("testdata/arrow_points_left.svg")
}

func (t *one_arrow) can_point_down() {
	t.End.X = t.Start.X
	t.End.Y = t.Start.Y + 50
	t.saveAs("testdata/arrow_points_down.svg")
}

func (t *one_arrow) can_point_up() {
	t.End.X = t.Start.X
	t.End.Y = t.Start.Y - 40
	t.saveAs("testdata/arrow_points_up.svg")
}

func (t *one_arrow) has_a_tail() {
	t.Tail = true
}

func (t *one_arrow) has_both_tail_and_head() {
	t.saveAs("testdata/arrow_with_tail_and_head.svg")
}

func (t *one_arrow) saveAs(filename string) {
	t.Helper()
	d := &Svg{Width: 100, Height: 100}
	d.Append(t.Arrow)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(style.NewStyler(fh))
	fh.Close()
}

func (t *one_arrow) can_have_a_specific_class() {
	t.Helper()
	t.Class = "special"
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	t.assert().Contains(buf.String(), "special")
}

func (t *one_arrow) can_move() {
	t.Helper()
	x, y := t.Position()
	t.SetX(x + 1)
	t.SetY(y + 1)
	t.assert(x != t.Start.X).Error("start X still the same")
	t.assert(y != t.Start.Y).Error("start Y still the same")
}

func (t *one_arrow) is_visible() {
	t.Helper()
	h := t.Height()
	w := t.Width()
	t.assert(h > 0 || w > 0).Errorf("%v not visible", t.Arrow)
}
