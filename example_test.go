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
	)
	diagram.Place(record).At(20, 20)
	diagram.Place(shapeI).RightOf(record, 90)
	diagram.Place(arrow).RightOf(shapeI, 90)
	diagram.Place(sws).Below(shapeI, 70)

	diagram.HAlignTop(record, shapeI, arrow)
	diagram.VAlignCenter(shapeI, sws)
	diagram.HAlignBottom(sws, arrow)
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
	)

	diagram.Place(diagramRec).At(10, 30)
	diagram.Place(record).RightOf(diagramRec)
	diagram.Place(adjuster).RightOf(record)
	diagram.Place(shapeI).Below(adjuster)

	diagram.HAlignTop(diagramRec, record, adjuster)
	diagram.HAlignCenter(record, diagramRec)
	diagram.HAlignBottom(record, shapeI)

	diagram.SaveAs("img/diagram_example.svg")
}

func TestClassDiagram(t *testing.T)    { ExampleClassDiagram() }
func TestSequenceDiagram(t *testing.T) { ExampleSequenceDiagram() }
func TestDiagram(t *testing.T)         { ExampleDiagram() }
