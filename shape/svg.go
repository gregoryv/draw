package shape

type Svg struct {
	Width, Height int
	Content       []svg
}

func (shape *Svg) Svg() string {
	xml := `<svg width="{{.Width}}" height="{{.Height}}">
{{ range .Content }}{{.Svg}}
{{end}}
</svg>`
	return toString(xml, shape)
}
