package design

import (
	"bytes"
	"io"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw/shape"
)

func TestDiagram(t *testing.T) {
	var (
		a = shape.NewRect("grid")
		b = shape.NewLabel("layout")
		c = shape.NewLabel("1")
		e = shape.NewCircle(30)
		g = shape.NewComponent("component")
	)

	t.Run("Adapts in size", func(t *testing.T) {
		d := NewDiagram()
		l1 := shape.NewLine(0, 0, 100, 100)
		l2 := shape.NewLine(0, 0, 100, 20)
		d.Place(l1).At(0, 0)
		d.Place(l2).Below(l1, 10)
		d.Append(&dummy{}) // Not a shape, should be skipped
		w, h := d.AdaptSize()
		assert := asserter.New(t)
		assert(w == 100).Errorf("width did not adapt: %v", w)
		assert(h == 130).Errorf("height did not adapt: %v", h)
	})

	t.Run("Can have fixed size", func(t *testing.T) {
		d := NewDiagram()
		d.Place(shape.NewLine(0, 0, 100, 100))
		adjusted := &bytes.Buffer{}
		d.WriteSvg(adjusted)

		d.SetWidth(5)
		d.SetHeight(10)
		fixed := &bytes.Buffer{}
		d.WriteSvg(fixed)

		assert := asserter.New(t)
		assert(adjusted.String() != fixed.String())
	})

	t.Run("Shapes can be placed in a grid", func(t *testing.T) {
		d := NewDiagram()
		cols := 2
		d.PlaceGrid(
			cols, 50, 20,
			a, b, e, c, g,
		)
		d.AdaptSize()
		d.SetHeight(d.Height() + 10)
		d.SetWidth(d.Width() + 10)
		d.SaveAs("img/grid_layout.svg")
	})
}

type dummy struct{}

func (d *dummy) WriteSvg(w io.Writer) error {
	_, err := w.Write([]byte("..."))
	return err
}
