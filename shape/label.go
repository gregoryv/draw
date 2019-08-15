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

func (shape *Label) Height() int {
	fontHeight := 10
	return fontHeight
}

func (label *Label) Width() int {
	fontWidth := 10
	return len(label.Text) * fontWidth
}

func (shape *Label) SetX(x int) { shape.X = x }
func (shape *Label) SetY(y int) { shape.Y = y }

func (shape *Label) Position() (int, int) {
	return shape.X, shape.Y
}
