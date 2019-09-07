package design

import (
	"os"
	"testing"

	"github.com/gregoryv/go-design/shape"
)

func new_one_arrow(t *testing.T) *one_arrow {
	return &one_arrow{
		T: t,
	}
}

type one_arrow struct {
	*testing.T
	shape.Arrow
}

func (t *one_arrow) starts_at_visible_position_N() {
	t.Helper()
	t.X1 = 50
	t.Y1 = 50
}

func (t *one_arrow) ends_above_and_to_the_right_of_N() {
	t.Helper()
	t.X2 = t.X1 + 30
	t.Y2 = t.Y1 - 30
}

func (t *one_arrow) points_up_and_to_the_right() {
	t.Helper()
	d := NewDiagram()
	t.Error(d.Width)
	d.SetWidth(150)
	t.Error(d.Width)
	d.SetHeight(150)
	d.Place(&t.Arrow)
	fh, err := os.Create("img/arrow_points_up_and_to_the_right.svg")
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(NewStyler(fh))
	fh.Close()
}
