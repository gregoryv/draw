package design

import "github.com/gregoryv/draw/shape"

func NewActivityDiagram() *ActivityDiagram {
	return &ActivityDiagram{
		Diagram: NewDiagram(),
		Spacing: 40,
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
	case shape.DirectionRight, shape.DirectionLeft:
		d.Place(label).Above(lnk, 20)
		d.VAlignCenter(lnk, label)
		shape.Move(label, -4, 0)
	case shape.DirectionUp, shape.DirectionDown:
		d.Place(label).RightOf(lnk, 5)
	}
}

func (d *ActivityDiagram) Place(next ...shape.Shape) *shape.Adjuster {
	d.last = next[0]
	return d.Diagram.Place(next...)
}

func (d *ActivityDiagram) Then(label string, txt ...string) *shape.Adjuster {
	next := shape.NewState(label)
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next, txt...)
	d.last = next
	return adj
}

// Decide adds a diamond below the last activity
func (d *ActivityDiagram) Decide() *shape.Diamond {
	next := shape.NewDecision()
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing+next.Height()/2)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next)
	d.last = next
	return next
}

// Exit adds an ExitDot below the last activity
func (d *ActivityDiagram) Exit(txt ...string) *shape.ExitDot {
	next := shape.NewExitDot()
	adj := d.Diagram.Place(next)
	adj.Below(d.last, d.Spacing)
	d.VAlignCenter(d.last, next)
	d.Link(d.last, next, txt...)
	d.last = next
	return next
}

// If adds next state to the right with a label.
func (d *ActivityDiagram) If(after shape.Shape, txt string, next shape.Shape) *shape.Adjuster {
	adj := d.Diagram.Place(next)
	adj.RightOf(after, d.Font.TextWidth(txt)+d.Spacing)
	d.HAlignCenter(after, next)
	d.Link(after, next, txt)
	d.last = next
	return adj
}

func (d *ActivityDiagram) Start() *shape.Adjuster {
	start := shape.NewDot()
	d.last = start
	return d.Place(start)
}
