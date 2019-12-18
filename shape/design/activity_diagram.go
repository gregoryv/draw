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

// Link places an arrow with a optional label above it between the two shapes.
func (diagram *ActivityDiagram) Link(from, to shape.Shape, txt ...string) *shape.Arrow {
	lnk := shape.NewArrowBetween(from, to)
	diagram.Place(lnk)
	if len(txt) > 0 {
		lnk.SetClass("activity-arrow")
		label := shape.NewLabel(txt[0])
		diagram.Place(label).Above(lnk, 20)
		diagram.VAlignCenter(lnk, label)
	}
	return lnk
}
