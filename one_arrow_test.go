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
	t.Start.X = 50
	t.Start.Y = 50
}

func (t *one_arrow) ends_above_and_right_of_N() {
	t.End.X = t.Start.X + 30
	t.End.Y = t.Start.Y - 30
}

func (t *one_arrow) ends_above_and_left_of_N() {
	t.End.X = t.Start.X - 30
	t.End.Y = t.Start.Y - 30
}

func (t *one_arrow) ends_below_and_left_of_N() {
	t.End.X = t.Start.X - 10
	t.End.Y = t.Start.Y + 30
}

func (t *one_arrow) ends_below_and_right_of_N() {
	t.End.X = t.Start.X + 20
	t.End.Y = t.Start.Y + 30
}

func (t *one_arrow) points_up_and_right() {
	t.saveAs("img/arrow_points_up_and_right.svg")
}

func (t *one_arrow) points_up_and_left() {
	t.saveAs("img/arrow_points_up_and_left.svg")
}

func (t *one_arrow) points_down_and_left() {
	t.saveAs("img/arrow_points_down_and_left.svg")
}

func (t *one_arrow) points_down_and_right() {
	t.saveAs("img/arrow_points_down_and_right.svg")
}

func (t *one_arrow) saveAs(filename string) {
	t.Helper()
	d := NewDiagram()
	d.SetWidth(100)
	d.SetHeight(100)
	d.Place(&t.Arrow)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(NewStyler(fh))
	fh.Close()
}
