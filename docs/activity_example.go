package docs

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleActivityDiagram() *design.ActivityDiagram {
	d := design.NewActivityDiagram()
	d.Spacing = 50
	d.Start().At(80, 20)
	d.Then("Commited", "Push")
	d.Then("Build complete", "run git hook")

	test := d.Decide("run tests")
	d.Then("Verified", "ok")
	d.Exit("deploy")

	d.If(test, "failed", shape.NewExitDot())
	return d
}
