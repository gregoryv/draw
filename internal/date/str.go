package date

import (
	"fmt"
	"time"
)

// String has the format of yyyymmdd
type String string

func (s String) Time() time.Time {
	var (
		year  string
		month string
		day   string
	)
	switch len(s) {
	case 8:
		year = string(s[:4])
		month = string(s[4:6])
		day = string(s[6:])
	default:
		panic(fmt.Sprintf("unexpeced format yyyymmdd: %s", s))
	}
	str := fmt.Sprintf("%s-%02s-%02sT00:00:00.000Z", year, month, day)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}

func (s String) DaysAfter(t time.Time) int {
	dur := s.Time().Sub(t)
	return int(dur.Hours() / 24)
}

// todo datestr+span
