package docs

import (
	"github.com/gregoryv/draw/design"
)

func ExampleActivityDiagram() *design.ActivityDiagram {
	d := design.NewActivityDiagram()

	d.Start().At(80, 20)
	d.Trans("push", "Commited")
	d.Trans("run git hook", "Build complete")

	test := d.Decide("run tests")
	d.Trans("ok", "Verified")
	d.Trans("", "EXIT")

	d.Or(test)
	d.TransRight("fails", "EXIT")
	return d
}
