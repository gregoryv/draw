package design_test

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleActivityDiagram() {
	var (
		d = design.NewActivityDiagram()
	)
	d.Start().At(80, 20)
	d.Then("Push commit")
	d.Then("Run git hook")
	dec := d.Decide()
	d.Then("Deploy", "ok")
	d.Exit()
	d.If(dec, "Tests failed", shape.NewExitDot())
	// manual part
	var (
		start = shape.NewDot()
		push  = shape.NewState("Push tag")
		hook  = shape.NewState("Run git hook")
		exit  = shape.NewExitDot()
	)
	d.Place(start).At(180, 20)
	d.Place(push).RightOf(start)
	d.Place(hook, exit).Below(push)
	d.HAlignCenter(start, push)
	d.VAlignCenter(push, hook, exit)
	d.LinkAll(start, push, hook, exit)
	d.SaveAs("img/activity_diagram.svg")
}
