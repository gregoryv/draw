/*
Package draw provides SVG writing features.

	s := NewSvg()
	s.WriteSVG(os.Stdout)

	<svg
	  xmlns="http://www.w3.org/2000/svg"
	  xmlns:xlink="http://www.w3.org/1999/xlink"
	  width="100" height="100" font-family="Arial, Helvetica, sans-serif"></svg>

*/
package draw

import (
	"io"
)

// NewSvg returns an empty Svg of size 100x100
func NewSvg() *SVG {
	return &SVG{
		width:   100,
		height:  100,
		Content: make([]SVGWriter, 0),
	}
}

type SVG struct {
	width, height int
	Content       []SVGWriter
}

func (s *SVG) WriteSVG(out io.Writer) error {
	w, err := NewTagWriter(out)
	w.Printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v" font-family="Arial, Helvetica, sans-serif">`, s.width, s.height)

	for _, c := range s.Content {
		w.Print("\n")
		c.WriteSVG(w)
	}
	w.Print("</svg>")
	return *err
}

func (s *SVG) Append(w ...SVGWriter) {
	s.Content = append(s.Content, w...)
}

func (s *SVG) Prepend(w ...SVGWriter) {
	s.Content = append(w, s.Content...)
}

func (s *SVG) Width() int  { return s.width }
func (s *SVG) Height() int { return s.height }

func (s *SVG) SetWidth(w int)   { s.width = w }
func (s *SVG) SetHeight(h int)  { s.height = h }
func (s *SVG) SetSize(w, h int) { s.SetWidth(w); s.SetHeight(h) }

type SVGWriter interface {
	WriteSVG(io.Writer) error
}
