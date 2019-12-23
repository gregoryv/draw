package date

import (
	"testing"
	"time"
)

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

func expectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Expected panic")
	}
}
