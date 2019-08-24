package design

import (
	"github.com/gregoryv/go-design/shape"
)

func ExampleClassDiagram() {
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

	diagram.HAlignTop(diagramRec, record, adjuster)
	diagram.Place(shapeI).Below(adjuster)
	diagram.SaveAs("img/class_example.svg")
}
