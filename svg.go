// package draw provides svg writing features
package draw

import (
	"io"
)

// NewSvg returns an empty Svg of size 100x100
func NewSvg() *Svg {
	return &Svg{
		width:   100,
		height:  100,
		Content: make([]SvgWriter, 0),
	}
}

type Svg struct {
	width, height int
	Content       []SvgWriter
}

func (shape *Svg) WriteSvg(out io.Writer) error {
	w, err := NewTagPrinter(out)
	w.Printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v" font-family="Arial, Helvetica, sans-serif">`, shape.width, shape.height)

	for _, s := range shape.Content {
		w.Print("\n")
		s.WriteSvg(w)
	}
	w.Print("</svg>")
	return *err
}

func (svg *Svg) Append(shapes ...SvgWriter) {
	svg.Content = append(svg.Content, shapes...)
}

func (svg *Svg) Prepend(shapes ...SvgWriter) {
	svg.Content = append(shapes, svg.Content...)
}

func (svg *Svg) Width() int  { return svg.width }
func (svg *Svg) Height() int { return svg.height }

func (svg *Svg) SetWidth(w int)   { svg.width = w }
func (svg *Svg) SetHeight(h int)  { svg.height = h }
func (svg *Svg) SetSize(w, h int) { svg.SetWidth(w); svg.SetHeight(h) }

type SvgWriter interface {
	WriteSvg(io.Writer) error
}
