package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAdjusterAt(t *testing.T) {
	it := NewOneAdjuster(t)
	it.HasDefaultSpacing()
	it.TakesOptionalSpacing()
	it.CanPositionShapes()
}

func NewOneAdjuster(t *testing.T) *OneAdjuster {
	return &OneAdjuster{t, NewAdjuster(&Line{}), asserter.New(t)}
}

type OneAdjuster struct {
	*testing.T
	*Adjuster
	assert asserter.AssertFunc
}

func (t *OneAdjuster) HasDefaultSpacing() {
	s := t.space([]int{})
	t.assert(s == t.Spacing).Error("No default spacing")
}

func (t *OneAdjuster) TakesOptionalSpacing() {
	s := t.space([]int{1})
	t.assert(s == 1).Error("no optional spacing")
}

func (t *OneAdjuster) CanPositionShapes() {
	t.At(1, 1)
	s := t.shapes[0]
	x, y := s.Position()
	t.assert(x == 1).Errorf("%v", x)
	t.assert(y == 1).Errorf("%v", x)

	o := &Line{}
	t.RightOf(o)
	x, _ = s.Position()
	t.assert(x == 30).Errorf("%v", x)

	t.LeftOf(o)
	x, _ = s.Position()
	t.assert(x == -30).Errorf("%v", x)

	t.Below(o)
	_, y = s.Position()
	t.assert(y == 30).Errorf("%v", y)

	t.Above(o)
	_, y = s.Position()
	t.assert(y == -30).Errorf("%v", y)
}
