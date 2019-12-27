package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func TestGanttChart_WriteSvg(t *testing.T) {
	w := bytes.NewBufferString("")
	var (
		d   = NewGanttChartFrom(30, "20191111")
		dev = d.Add("Develop")
		rel = d.Add("Release").Red()
		vac = d.Add("Vacation").Blue()
	)
	d.MarkDate("20191120")
	d.Place(dev).At("20191111", 10)
	d.Place(rel).At("20191121", 1)
	d.Place(vac).At("20191125", 14)
	d.SetCaption("Figure 1. Project estimated delivery")
	d.WriteSvg(w)
	golden.Assert(t, w.String())
}

func TestNewGanttChart(t *testing.T) {
	NewGanttChartFrom(20, "20191002")
	NewGanttChartFrom(20, "20190101")
	NewGanttChartFrom(20, "20190228")
}

func TestNewGanttChartFrom_panics(t *testing.T) {
	defer expectPanic(t)
	NewGanttChartFrom(20, "201910-2")
}

func expectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Expected panic")
	}
}

func TestGanttChart_MarkDate(t *testing.T) {
	d := NewGanttChartFrom(20, "20191002")
	d.MarkDate("20191003")
	d.MarkDate("20191204") // Ok even if it's outside the visible span
}

func TestGanttChart_MarkDate_panics(t *testing.T) {
	defer expectPanic(t)
	d := NewGanttChartFrom(20, "20191002")
	d.MarkDate("")
}
