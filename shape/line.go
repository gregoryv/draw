package shape

import (
	"bytes"
	"text/template"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

func (line *Line) Svg() string {
	xml := `<line x1="{{.X1}}" y1="{{.Y1}}" x2="{{.X2}}" y2="{{.Y2}}"`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, line)
	return buf.String()
}
