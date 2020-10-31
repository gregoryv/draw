package docs

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleSmallClassDiagram() *design.ClassDiagram {
	var (
		d      = design.NewClassDiagram()
		part   = d.Interface((*Part)(nil))
		door   = d.Struct(Door{})
		window = d.Struct(Window{})
		house  = d.Struct(House{})
		note   = shape.NewNote(`Relations are automatically rendered`)
	)
	d.Place(part).At(20, 20)        // absolute positioning
	d.Place(door).RightOf(part, 70) // optional extra spacing
	d.Place(window).Below(door)
	d.Place(house).RightOf(door, 70)
	d.Place(note).Below(window, 20)
	d.VAlignLeft(part, note)
	d.SetCaption("Small example diagram")
	return d
}

type House struct {
	Frontdoor Door      // aggregation
	Windows   []*Window // composition
}

func (me *House) Rooms() int { return 0 }

type Door struct {
	Material string
}

func (me *Door) Materials() []string {
	return []string{}
}

type Window struct {
	Model string
}

func (me *Window) Materials() []string {
	return []string{}
}

type Part interface {
	Materials() []string
}
