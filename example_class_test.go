package design

import "github.com/gregoryv/go-design/shape"

func ExampleClassDiagram() {
	diagram := NewDiagram()
	// todo reintroduce the Component for toggling fields
	diagram.Place(shape.NewRecordOf(diagram)).At(10, 10)

	diagram.SaveAs("img/class_example.svg")
}
