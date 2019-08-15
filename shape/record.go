package shape

import (
	"fmt"
	"io"
)

type Record struct {
	X, Y          int
	Width, Height int
	Title         string
	Public        []string

	Font    Font
	Padding Padding
}

func (shape *Record) WriteSvg(w io.Writer) error {
	_, e1 := fmt.Fprintf(w,
		`<rect x="%v" y="%v" width="%v" height="%v"/>`,
		shape.X, shape.Y, shape.Width, shape.Height)

	e2 := shape.title().WriteSvg(w)
	return firstOf(e1, e2)
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

type Font struct {
	Height int
	Width  int
}

type Padding struct {
	Left, Top, Right, Bottom int
}

func firstOf(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
