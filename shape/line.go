package shape

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

func (line *Line) Svg() string {
	xml := `<line x1="{{.X1}}" y1="{{.Y1}}" x2="{{.X2}}" y2="{{.Y2}}"/>`
	return toString(xml, line)
}
