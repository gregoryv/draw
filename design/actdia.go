package design

import (
	"github.com/gregoryv/draw/shape"
)

func NewActivityDiagram() *ActivityDiagram {
	return &ActivityDiagram{
		Diagram: NewDiagram(),
		Spacing: 60,
	}
}

type ActivityDiagram struct {
	*Diagram
	last    shape.Shape
	Spacing int
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
func (d *ActivityDiagram) Link(from, to shape.Shape, action ...string) *shape.Arrow {
	lnk := shape.NewArrowBetween(from, to)
	lnk.SetClass("activity-arrow")
	d.Place(lnk)
	if len(action) > 0 {
		d.placeLabel(lnk, action[0])
	}
	return lnk
}

func (d *ActivityDiagram) placeLabel(lnk *shape.Arrow, action string) {
	label := shape.NewLabel(action)
	switch lnk.Direction() {
	case shape.DirectionRight, shape.DirectionLeft:
		d.Place(label).Above(lnk, 20)
		d.VAlignCenter(lnk, label)
		shape.Move(label, -4, 0)
	case shape.DirectionUp, shape.DirectionDown:
		d.Place(label).RightOf(lnk, 5)
		d.HAlignCenter(lnk, label)
		shape.Move(label, 0, -8) // todo fix HAlignCenter instead
	}
}

func (d *ActivityDiagram) Place(next ...shape.Shape) *shape.Adjuster {
	d.last = next[0]
	return d.Diagram.Place(next...)
}

// When adds a new state below the last one
func (d *ActivityDiagram) Trans(action, label string) *shape.Adjuster {
	next := shapeFromLabel(label)
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next, action)
	d.last = next
	return adj
}

// WhenRight adds a new right of last state
func (d *ActivityDiagram) TransRight(action, label string) *shape.Adjuster {
	next := shapeFromLabel(label)
	adj := d.Diagram.Place(next)
	w := d.Font.TextWidth(action) + d.Spacing
	adj.RightOf(d.last, w)
	d.HAlignCenter(d.last, next)
	d.Link(d.last, next, action)
	d.last = next
	return adj
}

func shapeFromLabel(label string) shape.Shape {
	if label == "EXIT" {
		return shape.NewExitDot()
	}
	return shape.NewState(label)
}

// Decide adds a diamond below the last activity
func (d *ActivityDiagram) Decide(action ...string) *shape.Diamond {
	next := shape.NewDecision()
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing+next.Height()/2)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next, action...)
	d.last = next
	return next
}

// Exit adds an ExitDot below the last activity
func (d *ActivityDiagram) Exit(action ...string) *shape.ExitDot {
	next := shape.NewExitDot()
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next, action...)
	d.last = next
	return next
}

func (d *ActivityDiagram) Or(next shape.Shape) {
	d.last = next
}

func (d *ActivityDiagram) Start() *shape.Adjuster {
	start := shape.NewDot()
	d.last = start
	return d.Place(start)
}
