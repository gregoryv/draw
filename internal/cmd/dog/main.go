package main

import (
	"database/sql"

	"github.com/gregoryv/draw/internal/app"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/draw/shape/design"
)

//go:generate go run .
func main() {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})
	)
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	d.Link(srv, srv, "Transform to view model").Class = "highlight"
	d.Link(srv, cli, "Send HTML")
	d.SaveAs("../../../shape/design/img/app_sequence_diagram.svg")

	genOverview()
}

func genOverview() {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})

		start = shape.NewDot()
		run   = shape.NewState("Run")
		end   = shape.NewExitDot()

		circle = design.NewStruct(shape.Circle{})
		shp    = design.NewInterface((*shape.Shape)(nil))
		note   = shape.NewNote("Anything is possible!\nGo draw your next design")

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
	d.SaveAs("../../../overview.svg")
}
