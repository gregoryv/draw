package draw_test

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/internal/app"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/golden"
)

func BenchmarkWriteSvg(b *testing.B) {
	svg := &draw.SVG{}
	svg.Append(&shape.Record{})

	style := draw.NewStyle(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		svg.WriteSVG(&style)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func ExampleNewSvg() {
	s := draw.NewSVG()
	s.WriteSVG(os.Stdout)
	// output:
	// <svg
	//   xmlns="http://www.w3.org/2000/svg"
	//   xmlns:xlink="http://www.w3.org/1999/xlink"
	//   class="root" width="100" height="100"></svg>
}

func Test_overview(t *testing.T) {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})

		start = shape.NewDot()
		run   = shape.NewState("Run")
		end   = shape.NewExitDot()

		circle = design.NewVRecord(shape.Circle{})
		cyl    = shape.NewCylinder(40, 70)

		shp  = design.NewVRecord((*shape.Shape)(nil))
		note = shape.NewNote("Anything is possible!\nGo draw your next design")

		actor = shape.NewActor()
	)
	circle.HideMethods()
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "...")

	d.Place(start).At(100, 140)
	d.Place(run, end).Below(start)

	d.VAlignCenter(start, run, end)
	d.LinkAll(start, run, end)

	d.Place(shp).RightOf(start, 100)
	d.Place(circle).RightOf(shp, 100)
	d.HAlignCenter(shp, circle)
	d.Place(cyl).Below(circle)

	lnk := shape.NewArrowBetween(circle, shp)
	lnk.SetClass("implements-arrow")
	d.Place(lnk)
	label := shape.NewLabel("implements")
	d.Place(label).Above(lnk, 20)
	d.VAlignCenter(lnk, label)

	d.Place(note).Above(circle)
	d.Place(actor).Above(note)
	d.VAlignRight(note, actor)

	d.Place(shape.NewArrowBetween(actor, note))

	var buf bytes.Buffer
	d.Style.SetOutput(&buf)
	d.WriteSVG(&d.Style)

	golden.AssertWith(t, buf.String(), "overview.svg")
}
