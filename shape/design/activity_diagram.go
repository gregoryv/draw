package design

import "github.com/gregoryv/draw/shape"

func NewActivityDiagram() *ActivityDiagram {
	return &ActivityDiagram{
		Diagram: NewDiagram(),
	}
}

type ActivityDiagram struct {
	Diagram
}

// LinkAll places arrows between each shape, s0->s1->...->sn
func (diagram *ActivityDiagram) LinkAll(s ...shape.Shape) {
	for i, next := range s[1:] {
		lnk := shape.NewArrowBetween(s[i], next)
		lnk.SetClass("activity-arrow")
		diagram.Place(lnk)
	}
}

// Link places an arrow with a label above it between the two shapes.
func (diagram *ActivityDiagram) Link(from, to shape.Shape, txt string) {
	lnk := shape.NewArrowBetween(from, to)
	diagram.Place(lnk)
	lnk.SetClass("activity-arrow")
	label := shape.NewLabel(txt)
	diagram.Place(label).Above(lnk, 20)
	diagram.VAlignCenter(lnk, label)
}
