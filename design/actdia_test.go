package design_test

import (
	"github.com/gregoryv/draw/design"
)

func ExampleActivityDiagram() {
	d := design.NewActivityDiagram()

	d.Start().At(80, 20)
	d.Trans("push", "Commited")
	d.Trans("run git hook", "Build complete")

	test := d.Decide("unit test")
	d.TransRight("fails", "EXIT")

	d.Or(test)
	d.Trans("ok", "Verified")

	d.Trans("deploy to stage", "Deployed")

	itest := d.Decide("integration test")
	d.TransRight("fails", "EXIT")

	d.Or(itest)
	d.Trans("ok", "Verified")

	d.SaveAs("img/activity_diagram.svg")
}
