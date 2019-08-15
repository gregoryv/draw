package shape

import (
	"bytes"
	"io"
	"text/template"
)

type Svg struct {
	Width, Height int
	Content       []StyledWriter
}

// Style returns attribute style="..."
func (*Svg) Style() string {
	return ""
}

func (shape *Svg) WriteTo(w io.Writer) (int64, error) {
	xml := `<svg width="{{.Width}}" height="{{.Height}}" 
                 {{.Style}}/>`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, shape)
	return buf.WriteTo(w)
}
