package design

import (
	"bytes"
	"io"
	"text/template"
)

type Label struct {
	Pos
	StyleGuide
	Text string
}

func (label *Label) Width() int { return widthOf(label.Text) }

func (label *Label) WriteTo(w io.Writer) (int, error) {
	xml := `<text x="{{.X}}" y="{{.Y}}" {{.StrokeS}}>{{.Text}}</text>`
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, label)
	n, err := buf.WriteTo(w)
	return int(n), err // todo switch to int6r
}

func (label *Label) String() string {
	buf := bytes.NewBufferString("")
	label.WriteTo(buf)
	return buf.String()
}
