package shape

import (
	"bytes"
	"io"
	"text/template"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

// Style returns attribute style="..."
func (line *Line) Style() string {
	return `style="stroke:black;stroke-width:1"`
}

func (line *Line) WriteTo(w io.Writer) (int64, error) {
	xml := `<line x1="{{.X}}" y1="{{.Y}}" 
                  x2="{{.X2}}" y2="{{.Y2}}" {{.Style}}/>
`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, line)
	return buf.WriteTo(w)
}
