package design_test

import (
	"testing"

	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/draw/shape/design"
)

func ExampleClassDiagram() {
	var (
		d        = design.NewClassDiagram()
		record   = d.Struct(shape.Record{})
		arrow    = d.Struct(shape.Arrow{})
		line     = d.Struct(shape.Line{})
		circle   = d.Struct(shape.Circle{})
		diaarrow = d.Struct(shape.Diamond{})
		triangle = d.Struct(shape.Triangle{})
		shapE    = d.Interface((*shape.Shape)(nil))
	)
	d.HideRealizations()

	var (
		fnt      = d.Struct(shape.Font{})
		style    = d.Struct(shape.Style{})
		seqdia   = d.Struct(design.SequenceDiagram{})
		classdia = d.Struct(design.ClassDiagram{})
		dia      = d.Struct(design.Diagram{})
		aligner  = d.Struct(shape.Aligner{})
		adj      = d.Struct(shape.Adjuster{})
		rel      = d.Struct(design.Relation{})
	)
	d.HideRealizations()

	d.Place(shapE).At(220, 20)
	d.Place(record).At(20, 120)
	d.Place(line).Below(shapE, 90)
	d.VAlignCenter(shapE, line)

	d.Place(arrow).RightOf(line, 90)
	d.Place(circle).RightOf(shapE, 280)
	d.Place(diaarrow).Below(circle)
	d.Place(triangle).Below(diaarrow)
	d.HAlignBottom(record, arrow, line)

	d.Place(fnt).Below(record, 170)
	d.Place(style).RightOf(fnt, 90)
	d.VAlignCenter(shapE, line, style)
	d.VAlignCenter(record, fnt)

	d.Place(rel).Below(line, 80)
	d.Place(dia).RightOf(style, 90)
	d.Place(aligner).RightOf(dia, 80)
	d.HAlignCenter(fnt, style, dia, aligner)

	d.Place(adj).Below(fnt, 70)
	d.Place(seqdia).Below(aligner, 90)
	d.Place(classdia).Below(dia, 90)
	d.VAlignCenter(dia, classdia)
	d.HAlignBottom(classdia, seqdia)

	d.SetCaption("Figure 1. Class diagram of design and design.shape packages")
	d.SaveAs("img/class_example.svg")
}

func ExampleDiagram() {
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
	for _, arrow := range []*shape.Arrow{
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
	d.SaveAs("img/diagram_example.svg")
}

func ExampleActivityDiagram() {
	var (
		d       = design.NewActivityDiagram()
		start   = shape.NewDot()
		push    = shape.NewState("Push commit")
		dec     = shape.NewDecision()
		run     = shape.NewState("Run git hook")
		deploy  = shape.NewState("Deploy")
		endOk   = shape.NewExitDot()
		endFail = shape.NewExitDot()
	)
	d.Place(start).At(80, 20)
	d.Place(push, run, dec, deploy, endOk).Below(start, 40)
	d.Place(endFail).RightOf(dec, 80)

	d.VAlignCenter(start, push, run, dec, deploy, endOk)
	d.HAlignCenter(dec, endFail)

	d.LinkAll(start, push, run, dec, deploy, endOk)
	d.Link(dec, endFail, "Tests failed")

	d.SaveAs("img/activity_diagram.svg")
}

func ExampleGanttChart() {
	var (
		d   = design.NewGanttChartFrom(30, "20191111")
		dev = d.Add("Develop")
		rel = d.Add("Release").Red()
		vac = d.Add("Vacation").Blue()
	)
	d.MarkDate("20191120")
	d.Place(dev).At("20191111", 10)
	d.Place(rel).At("20191121", 1)
	d.Place(vac).At("20191125", 14)
	d.SetCaption("Figure 1. Project estimated delivery")
	d.SaveAs("img/gantt_chart.svg")
}

func TestExamples(t *testing.T) {
	ExampleClassDiagram()
	//ExampleSequenceDiagram()
	ExampleDiagram()
	ExampleActivityDiagram()
	ExampleGanttChart()
}
