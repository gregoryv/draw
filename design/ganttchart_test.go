package design

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gregoryv/draw/internal/date"
	"github.com/gregoryv/golden"
)

func TestGanttAdjuster_At(t *testing.T) {
	task := NewTask("hepp")
	a := &GanttAdjuster{
		start: date.String("20191001").Time(), // ie. diagram start
		task:  task,
	}
	// start before start of diagram
	a.At("20190930", 10)
}

func TestGanttChart_WriteSvg(t *testing.T) {
	w := bytes.NewBufferString("")
	var (
		d   = NewGanttChart("20191111", 30)
		dev = d.Add("Develop")
		rel = d.Add("Release").Red()
		vac = d.Add("Vacation").Blue()
	)
	d.MarkDate("20191120")
	d.Place(dev).At("20191111", 10)
	d.Place(rel).After(dev, 1)
	d.Place(vac).At("20191125", 14)
	d.SetCaption("Figure 1. Project estimated delivery")
	d.WriteSVG(w)
	golden.Assert(t, w.String())
}

func TestGanttChart_Inline(t *testing.T) {
	got := newTestChart().Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
}

func TestGanttChart_Inline_same_as_stringer(t *testing.T) {
	got := newTestChart().Inline()
	exp := newTestChart().String()
	if got != exp {
		t.Errorf(
			"inline is \n%s\n\n and stringer is \n %s",
			got, exp,
		)
	}
}

func TestNewGanttChart(t *testing.T) {
	NewGanttChart("20191002", 20)
	NewGanttChart("20190101", 20)
	NewGanttChart("20190228", 20)
}

func TestNewGanttChartFrom_panics(t *testing.T) {
	defer expectPanic(t)
	NewGanttChart("201910-2", 20)
}

func TestGanttChart_MarkDate(t *testing.T) {
	d := NewGanttChart("20191002", 20)
	d.MarkDate("20191003")
	d.MarkDate("20191204") // Ok even if it's outside the visible span
}

func TestGanttChart_MarkDate_panics(t *testing.T) {
	defer expectPanic(t)
	d := NewGanttChart("20191002", 20)
	d.MarkDate("")
}

func newTestChart() *GanttChart {
	var (
		d   = NewGanttChart("20191111", 30)
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

func expectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Expected panic")
	}
}
