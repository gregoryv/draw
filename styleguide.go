package design

type StyleGuide struct {
	FontWidth     int
	LineHeight    int
	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
}

var DefaultStyle = &StyleGuide{
	FontWidth:     16,
	LineHeight:    12,
	PaddingTop:    4,
	PaddingBottom: 4,
	PaddingLeft:   16,
}

func widthOf(txt string) int {
	return len(txt) * DefaultStyle.FontWidth
}

func (s *StyleGuide) Height(lines int) int {
	h := 2 // below line, e.g. letter 'g'
	return lines * (h + s.LineHeight + s.PaddingTop + s.PaddingBottom)
}

func (s *StyleGuide) Offset(x, y int) *Offset {
	return &Offset{x, y, s}
}

type Offset struct {
	x, y int
	s    *StyleGuide
}

func (o *Offset) Line(n int) int {
	return o.y + (o.s.LineHeight+o.s.PaddingTop+o.s.PaddingBottom)*n
}
