package date

import (
	"fmt"
	"time"
)

// DateStr has the format of yyyymmdd
type DateStr string

func (s DateStr) Time() time.Time {
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

// todo datestr+span
