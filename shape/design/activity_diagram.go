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
func (d *ActivityDiagram) LinkAll(s ...shape.Shape) {
	for i, next := range s[1:] {
		lnk := shape.NewArrowBetween(s[i], next)
		lnk.SetClass("activity-arrow")
		d.Place(lnk)
	}
}

// Link places an arrow with a optional label above it between the two
// shapes.
func (d *ActivityDiagram) Link(from, to shape.Shape, txt ...string) *shape.Arrow {
	lnk := shape.NewArrowBetween(from, to)
	lnk.SetClass("activity-arrow")
	d.Place(lnk)
	if len(txt) > 0 {
		d.placeLabel(lnk, txt[0])
	}
	return lnk
}

func (d *ActivityDiagram) placeLabel(lnk *shape.Arrow, txt string) {
	label := shape.NewLabel(txt)
	switch lnk.Direction() {
	case shape.RightDir, shape.LeftDir:
		d.Place(label).Above(lnk, 20)
		d.VAlignCenter(lnk, label)
		shape.Move(label, -4, 0)
	case shape.Up, shape.Down:
		d.Place(label).RightOf(lnk, 5)
	}
}
