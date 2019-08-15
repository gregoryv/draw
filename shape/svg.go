package shape

import (
	"bytes"
	"text/template"
)

type Svg struct {
	Width, Height int
	Content       []svg
}

func (shape *Svg) Svg() string {
	buf := bytes.NewBufferString("")
	Templates.ExecuteTemplate(buf, "svg", shape)
	return buf.String()
}

var (
	Templates = template.Must(template.New("svg").Parse(
		`<svg width="{{.Width}}" height="{{.Height}}">
{{ range .Content }}{{.Svg}}
{{end}}
</svg>`))
)
