package design_test

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleActivityDiagram() {
	var (
		d = design.NewActivityDiagram()
	)
	d.Spacing = 60
	d.Start().At(80, 20)
	d.Then("Commited", "Push")
	d.Then("Build complete", "run git hook")
	dec := d.Decide("run tests")
	d.Then("Verified", "ok")
	d.Exit("deploy")
	d.If(dec, "failed", shape.NewExitDot())

	d.SaveAs("img/activity_diagram.svg")
}
