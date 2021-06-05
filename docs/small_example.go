package docs

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleSmallClassDiagram() *design.ClassDiagram {
	var (
		d      = design.NewClassDiagram()
		house  = d.Struct(House{})
		door   = d.Struct(Door{})
		window = d.Struct(Window{})
		part   = d.Interface((*Part)(nil))
		note   = shape.NewNote(`Relations are automa-
tically rendered`)
	)
	d.Place(part).At(20, 20)        // absolute positioning
	d.Place(door).RightOf(part, 70) // optional extra spacing
	d.Place(window).Below(door)
	d.Place(house, note).RightOf(door, 70)

	d.HAlignCenter(house, door, part, note)
	d.SetCaption("Small example diagram")
	return d
}
