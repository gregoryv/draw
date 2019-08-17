package shape

import (
	"io"
)

type Svg struct {
	Width, Height int
	Content       []SvgWriterShape
}

func (shape *Svg) WriteSvg(w io.Writer) error {
	w, printf, err := newTagPrinter(w)
	printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v">`, shape.Width, shape.Height)

	for _, s := range shape.Content {
		printf("\n")
		s.WriteSvg(w)
	}
	printf("</svg>")
	return *err
}
