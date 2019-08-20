package shape

import (
	"io"
)

type Svg struct {
	Width, Height int
	Content       []SvgWriterShape
}

func (shape *Svg) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v" font-family="Arial, Helvetica, sans-serif">`, shape.Width, shape.Height)

	for _, s := range shape.Content {
		w.print("\n")
		s.WriteSvg(w)
	}
	w.print("</svg>")
	return *err
}

func (svg *Svg) Append(shapes ...SvgWriterShape) {
	svg.Content = append(svg.Content, shapes...)
}
