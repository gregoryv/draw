package shape

import (
	"io"
)

type Svg struct {
	width, height int
	Content       []Shape
}

func (shape *Svg) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v" font-family="Arial, Helvetica, sans-serif">`, shape.width, shape.height)

	for _, s := range shape.Content {
		w.print("\n")
		s.WriteSvg(w)
	}
	w.print("</svg>")
	return *err
}

func (svg *Svg) Append(shapes ...Shape) {
	svg.Content = append(svg.Content, shapes...)
}

func (svg *Svg) Prepend(shapes ...Shape) {
	svg.Content = append(shapes, svg.Content...)
}

func (svg *Svg) Width() int  { return svg.width }
func (svg *Svg) Height() int { return svg.height }

func (svg *Svg) SetWidth(w int)   { svg.width = w }
func (svg *Svg) SetHeight(h int)  { svg.height = h }
func (svg *Svg) SetSize(w, h int) { svg.SetWidth(w); svg.SetHeight(h) }
