package design

type StyleGuide struct {
	FontWidth int
}

var DefaultStyle = &StyleGuide{
	FontWidth: 18,
}

func widthOf(txt string) int {
	return len(txt) * DefaultStyle.FontWidth
}
