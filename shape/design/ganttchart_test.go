package design

import (
	"bytes"
	"testing"
	"time"

	"github.com/gregoryv/golden"
)

func TestGanttChart_WriteSvg(t *testing.T) {
	w := bytes.NewBufferString("")
	var (
		d = NewGanttChartFrom(30, "20191111")
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

func TestDateStr(t *testing.T) {
	var v DateStr = "20191101"
	got := v.Time()
	if got.Year() != 2019 {
		t.Errorf("bad year %v", got.Year())
	}
	if got.Month() != time.November {
		t.Errorf("bad month %v", got.Month())
	}
	if got.Day() != 1 {
		t.Errorf("bad day %v", got.Day())
	}
}

func TestDateStr_bad_format(t *testing.T) {
	t.Run("", func(t *testing.T) {
		defer expectPanic(t)
		var v DateStr = "hello"
		v.Time()
	})
	t.Run("", func(t *testing.T) {
		defer expectPanic(t)
		var v DateStr = "20191199"
		v.Time()
	})
}
