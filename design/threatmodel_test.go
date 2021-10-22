package design_test

import (
	"math"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func Example_threatModel() {
	var (
		d = design.NewDiagram()
		a = shape.NewActor()
		b = Entity("Customer\nDetails UI")
		c = Entity("Identity\nProvider")
		e = Entity("Customer\nDetails BFF")
		f = Entity("Customer\nService")
	)
	d.Style.Spacing = 60
	/* how about
		 d.Layout(
		`
	      4
	1 2 3
	      5`*/

	d.Place(a).At(20, 100)
	d.Place(b, c, e).RightOf(a)
	shape.Move(e, 0, -50)
	d.Place(f).Below(e)

	d.LinkAll(a, b, c)
	d.Link(c, e)
	d.Link(c, f)

	b1 := Boundary(b, c)
	d.Place(b1)

	d.SetCaption("Figure 3. Threat model diagram")
	d.SaveAs("img/threatmodel_example.svg")
	// output:
}

// Boundary
func Boundary(s1, s2 shape.Shape) shape.Shape {
	x1, y1 := s1.Position()
	x2, y2 := s2.Position()

	xd := x2 - (x1 + s1.Width())
	x := x1 + s1.Width() + xd/2

	e := 10 // extra
	s := 3  // slant, a bit of an angle
	l := shape.NewLine(x+s, y1-e, x-s, y2+s2.Height()+e)
	l.SetClass("boundary")
	return l
}

func Entity(v string) shape.Shape {
	return shape.NewRect(v)
}

// Shapes follow e.g.
// https://docs.microsoft.com/en-us/learn/modules/tm-create-a-threat-model-using-foundational-data-flow-diagram-elements/1b-elements
//
// Workflow maps into
// https://martinfowler.com/articles/agile-threat-modelling.html
func intAbs(v int) int {
	return int(math.Abs(float64(v)))
}
