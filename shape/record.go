package shape

import (
	"bytes"
	"io"
	"text/template"
)

type Record struct {
	X, Y          int
	Width, Height int
	Title         string
	Public        []string

	Font    Font
	Padding Padding
}

var recordSvg = template.Must(template.New("").Parse(
	`<rect x="{{.X}}" y="{{.Y}}"
     width="{{.Width}}" height="{{.Height}}"/>
{{.TitleSvg}}

`))

func (shape *Record) Svg() string {
	buf := bytes.NewBufferString("")
	recordSvg.Execute(buf, shape)
	return buf.String()
}

func (shape *Record) WriteSvg(w io.Writer) {
	recordSvg.Execute(w, shape)
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
