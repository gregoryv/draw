package docs

import (
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleDiagram() *design.Diagram {
	var (
		record     = shape.NewRecord("Record")
		x, y       = 130, 80
		q1arrow    = shape.NewArrow(x, y, x+50, y-10)
		q2arrow    = shape.NewArrow(x, y, x-30, y-10)
		q3arrow    = shape.NewArrow(x, y, x-50, y+20)
		q4arrow    = shape.NewArrow(x, y, x+40, y+20)
		rightarrow = shape.NewArrow(x, y, x+90, y)
		leftarrow  = shape.NewArrow(x, y, x-50, y)
		uparrow    = shape.NewArrow(x, y, x, y-40)
		downarrow  = shape.NewArrow(x, y, x, y+40)
		label      = shape.NewLabel("Label")
		withtail   = shape.NewArrow(20, 100, 150, 100)
		diaarrow   = shape.NewArrow(20, 120, 150, 120)
		note       = shape.NewNote(`Notes support
multilines`)
		comp   = shape.NewComponent("database")
		srv    = shape.NewComponent("service")
		circle = shape.NewCircle(10)
		dot    = shape.NewDot()
		exit   = shape.NewExitDot()
		rect   = shape.NewRect("Rect")
		state  = shape.NewState("Waiting for go routine")
		d      = design.NewDiagram()
	)
	d.Place(record).At(10, 30)
	for _, arrow := range []*shape.Line{
		q1arrow, q2arrow, q3arrow, q4arrow,
		rightarrow, leftarrow,
		uparrow, downarrow,
	} {
		d.Place(arrow)
	}
	d.Place(label).RightOf(record, 150)
	withtail.Tail = shape.NewCircle(3)
	d.Place(withtail).At(20, 150)
	diaarrow.Tail = shape.NewDiamond()
	d.Place(diaarrow).Below(withtail)
	d.Place(note).Below(diaarrow)
	d.Place(circle).Below(note)
	d.Place(dot).RightOf(circle)
	d.Place(exit).RightOf(dot)
	d.HAlignCenter(circle, dot, exit)
	d.LinkAll(circle, dot, exit)
	d.Place(comp).RightOf(diaarrow)
	d.Place(rect).Below(circle)
	d.Place(state).RightOf(rect)
	d.Place(srv).Below(comp)
	d.VAlignCenter(comp, srv)
	d.Link(srv, comp, "")
	return d
}

func ExampleGanttChart() *design.GanttChart {
	var (
		d   = design.NewGanttChart("20191111", 30)
		dev = d.Add("Develop")
		rel = d.Add("Release").Red()
		vac = d.Add("Vacation").Blue()
	)
	d.MarkDate("20191120")
	d.Place(dev).At("20191111", 10)
	d.Place(rel).After(dev, 1)
	d.Place(vac).At("20191125", 14)
	d.SetCaption("Figure 1. Project estimated delivery")
	return d
}
