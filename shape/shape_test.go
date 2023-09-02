package shape

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func TestContainer(t *testing.T) {

	label := NewLabel("hello")
	label.SetX(40)
	label.SetY(40)

	rect := NewRect("something")
	rect.SetX(120)
	rect.SetY(120)

	container := NewContainer(
		NewLabel("CONTAINER\nnote"),
		label,
		rect,
	)
	img := draw.NewSVG()
	img.SetHeight(250)
	img.SetWidth(300)
	img.Append(container, label, rect)
	file := "img/Container.svg"
	w, err := os.Create(file)
	if err != nil {
		t.Fatal(err)
	}
	style := draw.NewStyle()
	style.SetOutput(w)
	img.WriteSVG(&style)
	w.Close()
}

func TestGroup(t *testing.T) {
	label := NewLabel("hello")
	label.SetX(40)
	label.SetY(40)

	rect := NewRect("something")
	rect.SetX(120)
	rect.SetY(120)

	group := NewGroup(
		label,
		rect,
	)
	Move(group, -20, -20)
	img := draw.NewSVG()
	img.SetHeight(200)
	img.SetWidth(250)
	img.Append(group, label, rect)
	file := "img/Group.svg"
	w, err := os.Create(file)
	if err != nil {
		t.Fatal(err)
	}
	style := draw.NewStyle()
	style.SetOutput(w)
	img.WriteSVG(&style)
	w.Close()
}

func TestSaveShapes(t *testing.T) {
	shapes := []Shape{
		NewHexagon("hexagon", 80, 50, 20),
		NewDiamond(),
		NewComponent("a"),
		NewRect("rect"),
		NewTriangle(),
		NewLabel("l"),
		NewLine(1, 1, 7, 7),
		NewExitDot(),
		NewDot(),
		NewCircle(24),
		NewCylinder(40, 80),
		NewDatabase("postgres"),
		NewState("Waiting for push"),
		NewDecision(),
		NewActor(),
		NewInternet(),
		NewAnchor(
			"https://github.com/gregoryv/draw",
			NewComponent("package draw"),
		),
		NewHidden(
			NewInternet(),
		),
	}
	for _, shape := range shapes {
		img := draw.NewSVG()
		img.Append(shape)
		file := fmt.Sprintf("img/%T.svg", shape)
		file = strings.Replace(file, "*shape.", "", 1)
		w, err := os.Create(file)
		if err != nil {
			t.Fatal(err)
		}
		style := draw.NewStyle()
		style.SetOutput(w)
		img.WriteSVG(&style)
		w.Close()
	}
}

func TestShapes(t *testing.T) {
	r := NewRecord("a")
	r.Fields = []string{"bb", "ccc"}
	r.Methods = []string{"dddd", "eeeee"}
	shapes := []Shape{
		NewHexagon("hexagon", 80, 50, 20),
		NewComponent("a"),
		NewRect("rect"),
		NewTriangle(),
		NewLabel("l"),
		NewLine(1, 1, 7, 7),
		NewExitDot(),
		NewDot(),
		NewCircle(24),
		NewCylinder(40, 80),
		NewDatabase("postgres"),
		NewState("Waiting for push"),
		NewDecision(),
		NewActor(),
		NewGroup(NewLabel("hello")),
		NewContainer(
			NewLabel("hello"),
			NewCircle(24),
		),
		r,
	}
	for _, shape := range shapes {
		testShape(t, shape)
	}
	SetClass("x", shapes...)
}

func testShape(t *testing.T, shape Shape) {
	t.Helper()
	t.Run("Can move in X,Y direction", func(t *testing.T) {
		x, y := shape.Position()
		Move(shape, 1, 1)
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
		err := shape.WriteSVG(ioutil.Discard)
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
		s.Edge(xy.Point{0, 0})
	})

	t.Run("Is resizable", func(t *testing.T) {
		s, ok := shape.(resizable)
		if !ok {
			//t.Errorf("%T", shape)
			return
		}
		s.SetHeight(100)
		h := s.Height()
		if h != 100 {
			t.Fail()
		}
		s.SetWidth(100)
		w := s.Width()
		if w != 100 {
			t.Fail()
		}
	})

	t.Run("Special", func(t *testing.T) {
		s, ok := shape.(*Actor)
		if !ok {
			//t.Errorf("%T", shape)
			return
		}
		s.SetHeight(100)
		h := s.Height()
		if h != 100 {
			t.Fail()
		}
	})

}

type resizable interface {
	SetHeight(int)
	Height() int
	SetWidth(int)
	Width() int
}
