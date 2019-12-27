package date

import (
	"testing"
	"time"
)

func TestString_DaysAfter(t *testing.T) {
	a := String("20010101")
	b := String("20010102")
	got := b.DaysAfter(a.Time())
	exp := 1
	if got != exp {
		t.Errorf("%s-%s got days %v, expected %v", a, b, got, exp)
	}
}

func TestString(t *testing.T) {
	var v String = "20191101"
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

func TestString_bad_format(t *testing.T) {
	t.Run("", func(t *testing.T) {
		defer expectPanic(t)
		var v String = "hello"
		v.Time()
	})
	t.Run("", func(t *testing.T) {
		defer expectPanic(t)
		var v String = "20191199"
		v.Time()
	})
}

func expectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Expected panic")
	}
}
