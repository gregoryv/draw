package shape

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw"
)

func TestArrow_Height(t *testing.T) {
	// start and stop is the same, but head adds a height
	a := NewArrow(50, 50, 50, 50)
	if got := a.Height(); got == 0 {
		t.Error("unexpected height", got)
	}
}

func Test_arrowDirections(t *testing.T) {
	arrows := map[string]struct {
		*Line
		Direction
	}{
		"testdata/arrow_points_up_and_right.svg": {
			NewArrow(50, 50, 80, 20),
			DirectionUpRight,
		},
		"testdata/arrow_points_up_and_left.svg": {
			NewArrow(50, 50, 20, 20),
			DirectionUpLeft,
		},
		"testdata/arrow_points_down_and_left.svg": {
			NewArrow(50, 50, 40, 80),
			DirectionDownLeft,
		},
		"testdata/arrow_points_down_and_right.svg": {
			NewArrow(50, 50, 70, 80),
			DirectionDownRight,
		},
		"testdata/arrow_points_right.svg": {
			NewArrow(50, 50, 100, 50),
			DirectionRight,
		},
		"testdata/arrow_points_left.svg": {
			NewArrow(50, 50, 10, 50),
			DirectionLeft,
		},
		"testdata/arrow_points_down.svg": {
			NewArrow(50, 50, 50, 100),
			DirectionDown,
		},
		"testdata/arrow_points_up.svg": {
			NewArrow(50, 50, 50, 10),
			DirectionUp,
		},
	}

	for file, c := range arrows {
		saveAs(t, c.Line, file)
		dir := c.Line.Direction()
		if dir != c.Direction {
			t.Errorf("Direction: %v", dir)
		}
	}
}

func TestOneArrow(t *testing.T) {
	it := NewOneArrow(t)
	// when
	it.HasATail()
	it.HasBothTailAndHead()

	it.CanHaveASpecificClass()
	it.CanMove()
	it.IsVisible()
}

func NewOneArrow(t *testing.T) *OneArrow {
	return &OneArrow{t, NewArrow(50, 50, 50, 10), asserter.New(t)}
}

type OneArrow struct {
	*testing.T
	*Line
	assert
}

type assert = asserter.AssertFunc

func (t *OneArrow) HasATail() {
	t.Tail = NewCircle(3)
	t.Head = nil
	t.saveAs("testdata/arrow_with_tail.svg")
}

func (t *OneArrow) HasBothTailAndHead() {
	t.Tail = NewDiamond()
	t.Tail.SetX(20)
	t.Tail.SetY(24)
	t.Head = NewDiamond()
	t.Head.SetX(16)
	t.Head.SetY(8)
	t.saveAs("testdata/arrow_with_tail_and_head.svg")
}

func saveAs(t *testing.T, line *Line, filename string) {
	t.Helper()
	d := &draw.SVG{}
	d.SetSize(100, 100)
	d.Append(line)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	style := draw.NewStyle(fh)
	d.WriteSVG(&style)
	fh.Close()
}

func (t *OneArrow) saveAs(filename string) {
	t.Helper()
	d := &draw.SVG{}
	d.SetSize(100, 100)
	d.Append(t.Line)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	style := draw.NewStyle(fh)
	d.WriteSVG(&style)
	fh.Close()
}

func (t *OneArrow) CanHaveASpecificClass() {
	t.Helper()
	t.class = "special"
	buf := &bytes.Buffer{}
	t.WriteSVG(buf)
	t.assert().Contains(buf.String(), "special")
}

func (t *OneArrow) CanMove() {
	t.Helper()
	x, y := t.Position()
	t.SetX(x + 1)
	t.SetY(y + 1)
	t.assert(x != t.Start.X).Error("start X still the same")
	t.assert(y != t.Start.Y).Error("start Y still the same")
}

func (t *OneArrow) IsVisible() {
	t.Helper()
	h := t.Height()
	w := t.Width()
	t.assert(h > 0 || w > 0).Errorf("%v not visible", t.Line)
}

func TestLine(t *testing.T) {
	testShape(t, NewLine(0, 0, 50, 50))
}

func TestArrow_AbsAngle(t *testing.T) {
	a := NewLine(0, 0, 10, 10)
	got := a.AbsAngle()
	assert := asserter.New(t)
	assert().Equals(got, 45)
}

func TestArrow_Angle(t *testing.T) {
	a := NewLine(0, 0, 10, 10)
	got := a.Angle()
	assert := asserter.New(t)
	assert().Equals(got, 45)
}

// ----------------------------------------

func TestArrowBetweenShapes(t *testing.T) {
	it := &ArrowBetweenShapes{
		T:      t,
		assert: asserter.New(t),
	}

	it.StartsAndEndsAtEdgeOfShapes()
}

type ArrowBetweenShapes struct {
	*testing.T
	assert
}

func (t *ArrowBetweenShapes) StartsAndEndsAtEdgeOfShapes() {
	a := NewRecord("A")
	a.SetX(10)
	a.SetY(100)
	b := NewRecord("B")
	b.SetX(80)
	b.SetY(40)

	arrow := NewArrowBetween(a, b)
	svg := newSvg(200, 200, arrow, a, b)
	label := NewLabel(fmt.Sprintf("Angle: %v", arrow.absAngle()))
	label.SetX(100)
	label.SetY(80)

	svg.Append(label)
	writeSvgTo(t.T, "testdata/arrow_between_shapes.svg", svg)
}

func newSvg(width, height int, shapes ...draw.SVGWriter) *draw.SVG {
	svg := &draw.SVG{}
	svg.SetSize(width, height)
	svg.Append(shapes...)
	return svg
}

func writeSvgTo(t *testing.T, filename string, svg *draw.SVG) {
	t.Helper()
	fh, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	style := draw.NewStyle(fh)
	svg.WriteSVG(&style)
	fh.Close()
}
