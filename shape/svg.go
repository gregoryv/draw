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
	xml := `<svg width="{{.Width}}" height="{{.Height}}">
{{ range .Content }}{{.Svg}}{{end}}
</svg>`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, shape)
	return buf.String()
}
