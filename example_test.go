package design_test

import (
	"testing"

	design "github.com/gregoryv/go-design"
	"github.com/gregoryv/go-design/shape"
)

func ExampleClassDiagram() {
	var (
		d        = design.NewClassDiagram()
		record   = d.Struct(shape.Record{})
		arrow    = d.Struct(shape.Arrow{})
		line     = d.Struct(shape.Line{})
		circle   = d.Struct(shape.Circle{})
		diamond  = d.Struct(shape.Diamond{})
		triangle = d.Struct(shape.Triangle{})
		shapE    = d.Interface((*shape.Shape)(nil))
	)
	d.HideRealizations()

	var (
		fnt      = d.Struct(shape.Font{})
		style    = d.Struct(shape.Style{})
		seqdia   = d.Struct(design.SequenceDiagram{})
		classdia = d.Struct(design.ClassDiagram{})
		dia      = d.Struct(design.Diagram{})
		aligner  = d.Struct(shape.Aligner{})
		adj      = d.Struct(shape.Adjuster{})
	)
	d.HideRealizations()

	d.Place(shapE).At(220, 20)
	d.Place(record).At(20, 120)
	d.Place(line).Below(shapE, 90)
	d.VAlignCenter(shapE, line)

	d.Place(arrow).RightOf(line, 90)
	d.Place(circle).RightOf(shapE, 280)
	d.Place(diamond).Below(circle)
	d.Place(triangle).Below(diamond)
	d.HAlignBottom(record, arrow, line)

	d.Place(fnt).Below(record, 120)
	d.Place(style).RightOf(fnt, 90)
	d.VAlignCenter(shapE, line, style)
	d.VAlignCenter(record, fnt)

	d.Place(dia).RightOf(style, 90)
	d.Place(aligner).RightOf(dia, 80)
	d.HAlignCenter(fnt, style, dia, aligner)

	d.Place(adj).Below(fnt, 70)
	d.Place(seqdia).Below(aligner, 90)
	d.Place(classdia).Below(dia, 90)
	d.VAlignCenter(dia, classdia)
	d.HAlignBottom(classdia, seqdia)

	d.SetCaption("Figure 1. Class diagram of design and design.shape packages")
	d.SaveAs("img/class_example.svg")
}

func ExampleSequenceDiagram() {
	var (
		cli = "Client"
		srv = "Server"
		db  = "Database"
		d   = design.NewSequenceDiagram()
	)
	d.AddColumns(cli, srv, db)
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	// Special link
	lnk := d.Link(srv, srv, "Transform to view model")
	lnk.Class = "highlight"
	d.Link(srv, cli, "Send HTML")
	d.SaveAs("img/sequence_example.svg")
}

func ExampleDiagram() {
	var (
		record     = shape.NewRecord("Record")
		x, y       = 130, 80
		q1arrow    = shape.NewArrow(x, y, x+50, y-10)
		q2arrow    = shape.NewArrow(x, y, x-30, y-10)
		q3arrow    = shape.NewArrow(x, y, x-50, y+20)
		q4arrow    = shape.NewArrow(x, y, x+40, y+20)
		rightarrow = shape.NewArrow(x, y, x+90, y)
		leftarrow  = shape.NewArrow(x, y, x-50, y)
		uparrow    = shape.NewArrow(x, y, x, y-40)
		downarrow  = shape.NewArrow(x, y, x, y+40)
		label      = shape.NewLabel("Label")
		withtail   = shape.NewArrow(20, 100, 150, 100)
		diamond    = shape.NewArrow(20, 120, 150, 120)
		note       = shape.NewNote(`Notes support
multilines`)
		d = design.NewDiagram()
	)
	d.Place(record).At(10, 30)
	for _, arrow := range []*shape.Arrow{
		q1arrow, q2arrow, q3arrow, q4arrow,
		rightarrow, leftarrow,
		uparrow, downarrow,
	} {
		d.Place(arrow)
	}
	d.Place(label).RightOf(record, 150)
	withtail.Tail = shape.NewCircle(3)
	d.Place(withtail).At(20, 150)
	diamond.Tail = shape.NewDiamond(0, 0)
	d.Place(diamond).Below(withtail)
	d.Place(note).Below(diamond)
	d.SaveAs("img/diagram_example.svg")
}

func TestClassDiagram(t *testing.T) {
	ExampleClassDiagram()
}

func TestSequenceDiagram(t *testing.T) {
	ExampleSequenceDiagram()
}

func TestDiagram(t *testing.T) {
	ExampleDiagram()
}
