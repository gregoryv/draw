package shape

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gregoryv/asserter"
)

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
}
