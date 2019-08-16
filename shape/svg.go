package shape

import (
	"io"
)

type Svg struct {
	Width, Height int
	Content       []SvgWriter
}

func (shape *Svg) WriteSvg(w io.Writer) error {
	w, printf, err := newTagPrinter(w)
	printf(`<svg width="%v" height="%v">`, shape.Width, shape.Height)

	for _, s := range shape.Content {
		printf("\n")
		s.WriteSvg(w)
	}
	printf("</svg>")
	return *err
}
