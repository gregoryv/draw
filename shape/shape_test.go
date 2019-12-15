package shape

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw/xy"
)

func TestShapes(t *testing.T) {
	shapes := []Shape{
		NewComponent("a"),
		NewRect("rect"),
		NewTriangle(0, 0, ""),
		NewLabel("l"),
		NewLine(1, 1, 7, 7),
		NewExitDot(),
		NewDot(),
		NewCircle(24),
		NewState("Waiting for push"),
		NewDecision(),
		NewActor(),
	}
	for _, shape := range shapes {
		testShape(t, shape)
	}
}

func testShape(t *testing.T, shape Shape) {
	t.Helper()
	t.Run("Can move in X,Y direction", func(t *testing.T) {
		x, y := shape.Position()
		shape.SetX(x + 1)
		shape.SetY(y + 1)
		assert := asserter.New(t)
		x1, y1 := shape.Position()
		assert(x != x1).Errorf("x same as before move")
		assert(y != y1).Errorf("y same as before move")
	})

	t.Run("Has direction", func(t *testing.T) {
		assert := asserter.New(t)
		got := shape.Direction()
		assert(got >= 0).Errorf("Unknown direction: %v", got)
	})

	t.Run("Has width", func(t *testing.T) {
		assert := asserter.New(t)
		got := shape.Width()
		assert(got > 0).Errorf("0 width for: %v", shape)
	})

	t.Run("Has height", func(t *testing.T) {
		assert := asserter.New(t)
		got := shape.Height()
		assert(got >= 0).Errorf("0 height for: %v", shape)
	})

	t.Run("Implements fmt.Stringer", func(t *testing.T) {
		assert := asserter.New(t)
		s, ok := shape.(fmt.Stringer)
		assert(ok).Fatalf("%v", shape)
		got := s.String()
		assert(got != "").Error("String() returned empty string")
	})

	t.Run("May have class", func(t *testing.T) {
		shape.SetClass("something")
	})

	t.Run("Can be written as SVG", func(t *testing.T) {
		err := shape.WriteSvg(ioutil.Discard)
		assert := asserter.New(t)
		assert(err == nil).Error(err)
	})

	t.Run("Uses font", func(t *testing.T) {
		s, ok := shape.(HasFont)
		if !ok {
			return
		}
		s.SetFont(DefaultFont)
	})

	t.Run("Uses text padding", func(t *testing.T) {
		s, ok := shape.(HasTextPad)
		if !ok {
			return
		}
		s.SetTextPad(DefaultTextPad)
	})

	t.Run("Has edge", func(t *testing.T) {
		s, ok := shape.(Edge)
		if !ok {
			// TODO all shapes should have an edge so we can link
			// everything. Figure out what en edge means for lines and arrows.
			// t.Errorf("%T", shape)
			return
		}
		s.Edge(xy.Position{0, 0})
	})

	t.Run("Is resizable", func(t *testing.T) {
		s, ok := shape.(resizable)
		if !ok {
			//t.Errorf("%T", shape)
			return
		}
		s.SetHeight(100)
	})
}

type resizable interface {
	SetHeight(int)
}
