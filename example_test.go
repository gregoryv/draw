package design

import (
	"testing"

	"github.com/gregoryv/go-design/shape"
)

func ExampleClassDiagram() {
	diagram := NewClassDiagram()

	// todo
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

func TestSequenceDiagram(t *testing.T) { ExampleSequenceDiagram() }
func TestDiagram(t *testing.T)         { ExampleDiagram() }
