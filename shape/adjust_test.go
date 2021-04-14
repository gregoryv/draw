package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAdjusterAt(t *testing.T) {
	it := NewOneAdjuster(t)
	it.HasDefaultSpacing()
	it.takes_optional_spacing()
	it.can_position_shapes()
}

func NewOneAdjuster(t *testing.T) *OneAdjuster {
	return &OneAdjuster{t, NewAdjuster(&Line{})}
}

type OneAdjuster struct {
	*testing.T
	*Adjuster
}

func (t *OneAdjuster) HasDefaultSpacing() {
	s := t.space([]int{})
	if s != t.Spacing {
		t.Error("No default spacing")
	}
}

func (t *OneAdjuster) takes_optional_spacing() {
	s := t.space([]int{1})
	if s != 1 {
		t.Error("no optional spacing")
	}
}

func (t *OneAdjuster) can_position_shapes() {
	t.At(1, 1)
	s := t.shapes[0]
	x, y := s.Position()
	assert := asserter.New(t)
	assert(x == 1).Errorf("%v", x)
	assert(y == 1).Errorf("%v", x)

	o := &Line{}
	t.RightOf(o)
	x, _ = s.Position()
	assert(x == 30).Errorf("%v", x)

	t.LeftOf(o)
	x, _ = s.Position()
	assert(x == -30).Errorf("%v", x)

	t.Below(o)
	_, y = s.Position()
	assert(y == 30).Errorf("%v", y)

	t.Above(o)
	_, y = s.Position()
	assert(y == -30).Errorf("%v", y)
}
