package draw_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func TestInline(t *testing.T) {
	d := design.NewSequenceDiagram()
	got := d.Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attribute\n", got)
	}

}

func ExampleInline() {
	d := design.NewSequenceDiagram()
	fmt.Println(d.Inline())
	// output:
	// <svg
	//   xmlns="http://www.w3.org/2000/svg"
	//   xmlns:xlink="http://www.w3.org/1999/xlink"
	//   font-family="Arial,Helvetica,sans-serif" width="1" height="1"></svg>
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
	)
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "...")

	// Flow part
	var (
		start = shape.NewDot()
		run   = shape.NewState("Run")
		end   = shape.NewExitDot()
	)
	d.Place(start).At(100, 140)
	d.Place(run, end).Below(start)
	d.VAlignCenter(start, run, end)
	d.LinkAll(start, run, end)

	// Class diagram (manual)
	var (
		circle = design.NewVRecord(shape.Circle{})
		shp    = design.NewVRecord((*shape.Shape)(nil))
	)
	circle.HideMethods()
	d.Place(shp).RightOf(start, 100)
	d.Place(circle).RightOf(shp, 100)
	d.HAlignBottom(shp, circle)
	lnk := shape.NewArrowBetween(circle, shp)
	lnk.SetClass("implements-arrow")
	d.Place(lnk)

	// Actors and notes
	var (
		note  = shape.NewNote("Anything is possible!\nGo draw your next design")
		actor = shape.NewActor()
	)
	d.Place(note).Above(circle, 60)
	shape.Move(note, 50, 0)
	d.Place(actor).Above(note)
	d.VAlignLeft(note, actor)
	d.Place(shape.NewArrowBetween(actor, note))

	// components
	var (
		dbcomp  = shape.NewDatabase("database")
		inet    = shape.NewInternet()
		service = shape.NewComponent("Service")
		browser = shape.NewComponent("Browser")
	)
	browser.SetClass("external")
	d.Place(service).RightOf(note, 70)
	shape.Move(service, 0, -70)
	d.Place(dbcomp).RightOf(service)
	d.Place(inet).Below(service)
	d.Place(browser).Below(inet)
	d.VAlignCenter(service, inet, browser)

	d.SetCaption("gregoryv/draw provided shapes and diagrams")
	// Write it out inlined
	var buf bytes.Buffer
	d.Style.SetOutput(&buf)
	d.WriteSVG(&d.Style)

	golden.AssertWith(t, buf.String(), "overview.svg")
}

// The overview is used as social preview in github.
// Transform to png with e.g. inkscape -z -w 890 -h 356 overview.svg -e overview.png
