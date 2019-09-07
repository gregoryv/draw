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
	t.Start.X = 50
	t.Start.Y = 50
}

func (t *one_arrow) ends_above_and_to_the_right_of_N() {
	t.Helper()
	t.End.X = t.Start.X + 30
	t.End.Y = t.Start.Y - 30
}

func (t *one_arrow) points_up_and_to_the_right() {
	t.Helper()
	d := NewDiagram()
	d.SetWidth(100)
	d.SetHeight(100)
	d.Place(&t.Arrow)

	fh, err := os.Create("img/arrow_points_up_and_to_the_right.svg")
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(NewStyler(fh))
	fh.Close()
}
