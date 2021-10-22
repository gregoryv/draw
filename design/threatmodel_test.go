package design_test

import (
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
	d.Place(a).At(10, 10)
	d.Place(b, c, e).RightOf(a)
	d.Place(f).Below(e)

	d.SetCaption("Figure 3. Threat model diagram")
	d.SaveAs("img/threatmodel_example.svg")
	// output:
}

func Entity(v string) shape.Shape {
	return shape.NewRect(v)
}

// Shapes follow e.g.
// https://docs.microsoft.com/en-us/learn/modules/tm-create-a-threat-model-using-foundational-data-flow-diagram-elements/1b-elements
//
// Workflow maps into
// https://martinfowler.com/articles/agile-threat-modelling.html
