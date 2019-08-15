package shape

import (
	"bytes"
	"text/template"
)

type Label struct {
	X, Y int
	Text string
}

func (shape *Label) Svg() string {
	xml := `<text x="{{.X}}" y="{{.Y}}">{{.Text}}</text>`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, shape)
	return buf.String()
}
