package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func TestGanttChart_WriteSvg(t *testing.T) {
	w := bytes.NewBufferString("")
	var (
		d = NewGanttChartFrom(30, 2019, 11, 11)
	)
	d.MarkDate(2019, 11, 20)
	d.Add("Develop", 0, 10)
	d.Add("Release", 10, 1).Red()
	d.Add("Vacation", 14, 14).Blue()
	d.SetCaption("Figure 1. Project estimated delivery")
	d.WriteSvg(w)
	golden.Assert(t, w.String())
}

func TestNewGanttChart(t *testing.T) {
	NewGanttChartFrom(20, 2019, 10, 2)
}

func TestNewGanttChartFrom_panics(t *testing.T) {
	defer expectPanic(t)
	NewGanttChartFrom(20, 2019, 10, -2)
}

func expectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Expected panic")
	}
}

func TestGanttChart_MarkDate(t *testing.T) {
	d := NewGanttChartFrom(20, 2019, 10, 2)
	ok := func(err error) {
		t.Helper()
		if err != nil {
			t.Error(err)
		}
	}
	ok(d.MarkDate(2019, 10, 3))

	bad := func(err error) {
		t.Helper()
		if err == nil {
			t.Error("should fail")
		}
	}
	bad(d.MarkDate(-1, 0, 0))
}
