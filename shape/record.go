package shape

type Record struct {
	X, Y          int
	Width, Height int
	Title         string
	Public        []string
}

func (shape *Record) Svg() string {
	xml := `<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}"/>
{{.TitleSvg}}

`
	return toString(xml, shape)
}

func (record *Record) TitleSvg() string {
	fontHeight := 10
	label := &Label{
		Y:    record.Y + fontHeight,
		Text: record.Title,
	}
	return label.Svg()
}
