package design_test

import (
	"testing"

	design "github.com/gregoryv/go-design"
	"github.com/gregoryv/go-design/shape"
)

func ExampleClassDiagram() {
	d := design.NewClassDiagram()
	record := design.NewStruct(shape.Record{})
	d.Place(record).At(20, 20)
	shapeI := design.NewInterface((*shape.Shape)(nil))
	d.Place(shapeI).RightOf(record, 90)
	arrow := design.NewStruct(shape.Arrow{})
	d.Place(arrow).RightOf(shapeI, 90)
	sws := design.NewInterface((*shape.SvgWriterShape)(nil))
	d.Place(sws).Below(shapeI, 70)
	line := design.NewStruct(shape.Line{})
	d.Place(line).Below(arrow, 90)

	d.HAlignTop(record, shapeI)
	d.VAlignCenter(shapeI, sws)
	d.HAlignCenter(record, arrow)

	fnt := design.NewStruct(shape.Font{})
	d.Place(fnt).Below(record, 40)
	aligner := design.NewStruct(shape.Aligner{})
	d.Place(aligner).RightOf(fnt)

	d.HAlignTop(fnt, aligner)

	seqdia := design.NewStruct(design.SequenceDiagram{})
	d.Place(seqdia).Below(fnt, 100)

	d.VAlignLeft(fnt, seqdia)

	dia := design.NewStruct(design.Diagram{})
	d.Place(dia).RightOf(seqdia)

	d.SaveAs("img/class_example.svg")
}

func ExampleVerticalClassDiagram() {
	record := design.NewStruct(shape.Record{})
	record.TitleOnly()

	shapeI := design.NewInterface((*shape.Shape)(nil))
	shapeI.TitleOnly()

	sws := design.NewInterface((*shape.SvgWriterShape)(nil))
	sws.TitleOnly()

	arrow := design.NewStruct(shape.Arrow{})
	arrow.TitleOnly()

	d := design.NewClassDiagram()

	d.Place(shapeI).At(20, 100)
	d.Place(record).At(160, 0)
	d.Place(sws).At(280, 100)
	d.Place(arrow).At(60, 200)

	//d.HAlignTop(record, shapeI, arrow)
	d.VAlignCenter(record, arrow)
	//d.HAlignBottom(sws, arrow)

	d.SaveAs("img/vertical_class_example.svg")
}

func ExampleSequenceDiagram() {
	d := design.NewSequenceDiagram()
	cli, srv, db := "Client", "Server", "Database"
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
		diagramRec = shape.NewStructRecord(design.Diagram{})
		record     = shape.NewStructRecord(shape.Record{})
		adjuster   = shape.NewStructRecord(shape.Adjuster{})
		shapeI     = shape.NewInterfaceRecord((*shape.Shape)(nil))
		y          = 400
		q1arrow    = shape.NewArrow(230, y, 280, y-10)
		q2arrow    = shape.NewArrow(230, y, 200, y-10)
		q3arrow    = shape.NewArrow(230, y, 180, y+20)
		q4arrow    = shape.NewArrow(230, y, 270, y+20)
		rightarrow = shape.NewArrow(230, y, 320, y)
		leftarrow  = shape.NewArrow(230, y, 180, y)
		uparrow    = shape.NewArrow(230, y, 230, y-40)
		downarrow  = shape.NewArrow(230, y, 230, y+40)
		d          = design.NewDiagram()
	)
	d.Place(diagramRec).At(10, 30)
	d.Place(record).RightOf(diagramRec)
	d.Place(adjuster).RightOf(record)
	d.Place(shapeI).Below(adjuster)

	for _, arrow := range []*shape.Arrow{
		q1arrow, q2arrow, q3arrow, q4arrow,
		rightarrow, leftarrow,
		uparrow, downarrow,
	} {
		d.Place(arrow)
	}
	d.HAlignTop(diagramRec, record, adjuster)
	d.HAlignCenter(record, diagramRec)
	d.HAlignBottom(record, shapeI)

	d.SaveAs("img/diagram_example.svg")
}

func TestClassDiagram(t *testing.T)         { ExampleClassDiagram() }
func TestVerticalClassDiagram(t *testing.T) { ExampleVerticalClassDiagram() }
func TestSequenceDiagram(t *testing.T)      { ExampleSequenceDiagram() }
func TestDiagram(t *testing.T)              { ExampleDiagram() }
