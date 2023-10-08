package design_test

import (
	"testing"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/draw/xy"
)

func ExampleClassDiagram() *design.ClassDiagram {
	var (
		d        = design.NewClassDiagram()
		point    = d.Struct(xy.Point{})
		record   = d.Struct(shape.Record{})
		circle   = d.Struct(shape.Circle{})
		diamond  = d.Struct(shape.Diamond{})
		triangle = d.Struct(shape.Triangle{})
		shapE    = d.Interface((*shape.Shape)(nil))
	)
	d.HideRealizations()

	var (
		fnt      = d.Struct(draw.Font{})
		style    = d.Struct(draw.Style{})
		seqdia   = d.Struct(design.SequenceDiagram{})
		classdia = d.Struct(design.ClassDiagram{})
		dia      = d.Struct(design.Diagram{})
		aligner  = d.Struct(shape.Aligner{})
		adj      = d.Struct(shape.Adjuster{})
		rel      = d.Struct(design.Relation{})
	)
	d.HideRealizations()

	d.Place(shapE).At(280, 20)
	d.Place(record).At(20, 260)

	d.Place(diamond).RightOf(shapE, 200)
	d.Place(circle).Below(diamond, 50)
	d.Place(triangle).Below(circle)
	d.Place(point).RightOf(circle, 100)

	d.Place(fnt).Below(record, 120)
	d.Place(style).RightOf(fnt, 90)
	d.VAlignCenter(record, fnt)

	d.Place(rel).Below(shapE, 20)
	d.Place(dia).RightOf(style, 70)
	d.Place(adj).RightOf(dia, 60)
	d.Place(aligner).RightOf(dia, 60)
	d.HAlignCenter(fnt, style, dia, aligner)
	d.HAlignCenter(record, rel)
	d.HAlignTop(dia, adj)
	shape.Move(adj, 0, -60)

	d.Place(seqdia).Below(aligner, 90)
	d.Place(classdia).Below(dia, 90)
	d.VAlignCenter(dia, classdia)
	d.HAlignBottom(classdia, seqdia)

	d.HAlignCenter(circle, point)

	d.SetCaption("Figure 1. Class diagram of design and design.shape packages")
	d.SaveAs("img/class_example.svg")
	return d
}

func ExampleDiagram() {
	var (
		record     = shape.NewRecord("Record")
		x, y       = 130, 80
		q1arrow    = shape.NewLine(x, y, x+50, y-10)
		q2arrow    = shape.NewLine(x, y, x-30, y-10)
		q3arrow    = shape.NewLine(x, y, x-50, y+20)
		q4arrow    = shape.NewLine(x, y, x+40, y+20)
		rightarrow = shape.NewLine(x, y, x+90, y)
		leftarrow  = shape.NewLine(x, y, x-50, y)
		uparrow    = shape.NewLine(x, y, x, y-40)
		downarrow  = shape.NewLine(x, y, x, y+40)
		label      = shape.NewLabel("Label")
		withtail   = shape.NewLine(20, 100, 150, 100)
		diaarrow   = shape.NewLine(20, 120, 150, 120)
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

	// labeled arrows
	d1 := shape.NewDot()
	d.Place(d1).At(600, 200)

	d2 := shape.NewDot()
	d.Place(d2).RightOf(d1, 120)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).At(660, 260)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).Below(d1, 120)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).At(520, 280)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).LeftOf(d1, 120)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).At(520, 120)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).Above(d1, 120)
	d.Link(d1, d2, "label")

	d2 = shape.NewDot()
	d.Place(d2).At(660, 120)
	d.Link(d1, d2, "label")

	d.SaveAs("img/diagram_example.svg")
}

func ExampleGanttChart() {
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
	d.SaveAs("img/gantt_chart.svg")
}

func ExampleGanttChart_year() {
	var (
		d   = design.NewGanttChart("20191111", 365)
		dev = d.Add("Develop")
		rel = d.Add("Release").Red()
		vac = d.Add("Vacation").Blue()
	)
	d.Weeks = true
	d.MarkDate("20200404")
	d.Place(dev).At("20191111", 60)
	d.Place(rel).After(dev, 1)
	d.Place(vac).At("20200404", 21)
	d.SetCaption("Figure 1. Project estimated delivery")
	d.SaveAs("img/gantt_year.svg")
}

func TestExamples(t *testing.T) {
	ExampleClassDiagram()
	//ExampleSequenceDiagram()
	ExampleDiagram()
	ExampleActivityDiagram()
	ExampleGanttChart()
	ExampleGanttChart_year()
}
