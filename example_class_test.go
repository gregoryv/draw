package design

import "github.com/gregoryv/go-design/shape"

func ExampleClassDiagram() {
	diagram := NewDiagram()
	diagramRec := shape.NewRecordOf(diagram)
	record := shape.NewRecordOf(shape.Record{})

	diagram.Place(diagramRec).At(10, 30)
	diagram.Place(record).RightOf(diagramRec)

	shape.AlignHorizontal(shape.Center, diagramRec, record)
	diagram.SaveAs("img/class_example.svg")
}
