package design

import (
	"testing"

	"github.com/gregoryv/go-design/shape"
)

func ExampleClassDiagram() {
	var (
		diagram = NewClassDiagram()
		record  = NewStruct(shape.Record{})
		shapeI  = NewInterface((*shape.Shape)(nil))
		sws     = NewInterface((*shape.SvgWriterShape)(nil))
		arrow   = NewStruct(shape.Arrow{})
		fnt     = NewStruct(shape.Font{})
	)
	diagram.Place(record).At(20, 20)
	diagram.Place(shapeI).RightOf(record, 90)
	diagram.Place(arrow).RightOf(shapeI, 90)
	diagram.Place(sws).Below(shapeI, 70)

	diagram.HAlignTop(record, shapeI, arrow)
	diagram.VAlignCenter(shapeI, sws)
	diagram.HAlignBottom(sws, arrow)

	diagram.Place(fnt).Below(record, 40)

	diagram.SaveAs("img/class_example.svg")
}

func ExampleSequenceDiagram() {
	diagram := NewSequenceDiagram()
	cli, srv, db := "Client", "Server", "Database"
	diagram.AddColumns(cli, srv, db)
	diagram.Link(cli, srv, "connect()")
	diagram.Link(srv, db, "SELECT")
	diagram.Link(db, srv, "Rows")
	// Special link
	lnk := diagram.Link(srv, srv, "Transform to view model")
	lnk.Class = "highlight"
	diagram.Link(srv, cli, "Send HTML")

	diagram.SaveAs("img/sequence_example.svg")
}

func ExampleDiagram() {
	var (
		diagram    = NewDiagram()
		diagramRec = shape.NewStructRecord(Diagram{})
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
	)

	diagram.Place(diagramRec).At(10, 30)
	diagram.Place(record).RightOf(diagramRec)
	diagram.Place(adjuster).RightOf(record)
	diagram.Place(shapeI).Below(adjuster)

	for _, arrow := range []*shape.Arrow{
		q1arrow, q2arrow, q3arrow, q4arrow,
		rightarrow, leftarrow,
		uparrow, downarrow,
	} {
		diagram.Place(arrow)
	}

	diagram.HAlignTop(diagramRec, record, adjuster)
	diagram.HAlignCenter(record, diagramRec)
	diagram.HAlignBottom(record, shapeI)

	diagram.SaveAs("img/diagram_example.svg")
}

func TestClassDiagram(t *testing.T)    { ExampleClassDiagram() }
func TestSequenceDiagram(t *testing.T) { ExampleSequenceDiagram() }
func TestDiagram(t *testing.T)         { ExampleDiagram() }
