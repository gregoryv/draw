package shape

type Record struct {
	X, Y          int
	Width, Height int
	Title         string
	Public        []string

	Font    Font
	Padding Padding
}

func (shape *Record) Svg() string {
	xml := `<rect x="{{.X}}" y="{{.Y}}" width="{{.Width}}" height="{{.Height}}"/>
{{.TitleSvg}}

`
	return toString(xml, shape)
}

func (record *Record) TitleSvg() string {
	fontHeight := record.Font.Height
	padding := record.Padding.Left
	label := &Label{
		X:    record.X + padding,
		Y:    record.Y + fontHeight + padding,
		Text: record.Title,
	}
	return label.Svg()
}

type Font struct {
	Height int
	Width  int
}

type Padding struct {
	Left, Top, Right, Bottom int
}
