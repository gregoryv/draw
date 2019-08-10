package design

import "strconv"

type StyleGuide struct {
	FontWidth int
}

var DefaultStyle = &StyleGuide{
	FontWidth: 18,
}

func toFit(txt string) string {
	l := len(txt) * DefaultStyle.FontWidth
	return strconv.Itoa(l)
}
