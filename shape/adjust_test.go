package shape

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestAdjusterAt(t *testing.T) {
	it := new_one_adjuster(t)

	it.has_default_spacing()
	it.takes_optional_spacing()
	it.can_position_shapes()
}

func new_one_adjuster(t *testing.T) *one_adjuster {
	return &one_adjuster{t, NewAdjuster(&Line{})}
}

type one_adjuster struct {
	*testing.T
	*Adjuster
}

func (t *one_adjuster) has_default_spacing() {
	s := t.space([]int{})
	if s != t.defaultSpace {
		t.Error("No default spacing")
	}
}

func (t *one_adjuster) takes_optional_spacing() {
	s := t.space([]int{1})
	if s != 1 {
		t.Error("no optional spacing")
	}
}

func (t *one_adjuster) can_position_shapes() {
	t.At(1, 1)
	x, y := t.shape.Position()
	assert := asserter.New(t)
	assert(x == 1).Errorf("%v", x)
	assert(y == 1).Errorf("%v", x)

	o := &Line{}
	t.RightOf(o)
	x, _ = t.shape.Position()
	assert(x == 30).Errorf("%v", x)

	t.LeftOf(o)
	x, _ = t.shape.Position()
	assert(x == -30).Errorf("%v", x)

	t.Below(o)
	_, y = t.shape.Position()
	assert(y == 30).Errorf("%v", y)

	t.Above(o)
	_, y = t.shape.Position()
	assert(y == -30).Errorf("%v", y)
}
