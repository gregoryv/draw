// Document generation for example diagrams
package main

import (
	"database/sql"
	"flag"
	"path"

	"github.com/gregoryv/draw/internal/app"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/draw/shape/design"
)

//go:generate go run . ../../../img/
func main() {
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		root = "./"
	}
	overview().SaveAs(path.Join(root, "overview.svg"))
}

func overview() *design.SequenceDiagram {
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

		shp  = design.NewInterface((*shape.Shape)(nil))
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
	return d
}
