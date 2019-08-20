package design

func ExampleClassDiagram() {
	diagram := NewDiagram()
	// todo reintroduce the Component for toggling fields
	diagram.Place(*diagram).At(10, 10)

	diagram.SaveAs("img/class_example.svg")
}
