package shape

import (
	"fmt"
	"io"
)

type Record struct {
	X, Y          int
	Width, Height int
	Title         string
	PublicFields  []string

	Font    Font
	Padding Padding
}

func (shape *Record) WriteSvg(w io.Writer) error {
	collect := &ErrCollector{}
	collect.Last(fmt.Fprintf(w,
		`<rect x="%v" y="%v" width="%v" height="%v"/>`,
		shape.X, shape.Y, shape.Width, shape.Height))

	collect.Err(shape.title().WriteSvg(w))
	return collect.First()
}

func (record *Record) title() *Label {
	fontHeight := record.Font.Height
	padding := record.Padding.Left
	return &Label{
		X:    record.X + padding,
		Y:    record.Y + fontHeight + padding,
		Text: record.Title,
	}
}
