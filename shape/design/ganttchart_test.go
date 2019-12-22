package design

import "testing"

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
