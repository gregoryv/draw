package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw/shape"
)

func TestDiagram(t *testing.T) {
	d := NewDiagram()
	AdaptsInSize(t, &d)
	CanHaveFixedSize(t, &d)
}

func AdaptsInSize(t *testing.T, d *Diagram) {
	t.Helper()
	l1 := shape.NewLine(0, 0, 100, 100)
	l2 := shape.NewLine(0, 0, 100, 20)
	d.Place(l1).At(0, 0)
	d.Place(l2).Below(l1, 10)
	w, h := d.AdaptSize()
	assert := asserter.New(t)
	assert(w == 100).Errorf("width did not adapt: %v", w)
	assert(h == 130).Errorf("height did not adapt: %v", h)
}

func CanHaveFixedSize(t *testing.T, d *Diagram) {
	t.Helper()
	d.Place(shape.NewLine(0, 0, 100, 100))
	adjusted := &bytes.Buffer{}
	d.WriteSvg(adjusted)

	d.SetWidth(5)
	d.SetHeight(10)
	fixed := &bytes.Buffer{}
	d.WriteSvg(fixed)

	assert := asserter.New(t)
	assert(adjusted.String() != fixed.String())
}

func TestDiagram_PlaceGrid(t *testing.T) {
	var (
		d = NewDiagram()
		a = shape.NewRect("grid")
		b = shape.NewLabel("layout")
		c = shape.NewLabel("1")
		e = shape.NewCircle(30)
		g = shape.NewComponent("component")
	)
	cols := 2
	d.PlaceGrid(
		cols, 50, 20,
		a, b, e, c, g,
	)
	d.AdaptSize()
	d.SetHeight(d.Height() + 10)
	d.SetWidth(d.Width() + 10)
	d.SaveAs("img/grid_layout.svg")
}
