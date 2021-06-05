/*
Package draw provides SVG writing features.

	s := NewSVG()
	s.WriteSVG(os.Stdout)

	<svg
	  xmlns="http://www.w3.org/2000/svg"
	  xmlns:xlink="http://www.w3.org/1999/xlink"
	  width="100" height="100" font-family="Arial, Helvetica, sans-serif"></svg>

*/
package draw

import (
	"io"

	"github.com/gregoryv/nexus"
)

// NewSVG returns an empty SVG of size 100x100
func NewSVG() *SVG {
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
	w, err := nexus.NewPrinter(out)
	w.Printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  class="root" width="%v" height="%v">`, s.width, s.height)

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

func (s *SVG) SetWidth(w int)  { s.width = w }
func (s *SVG) SetHeight(h int) { s.height = h }
func (s *SVG) SetSize(width, height int) {
	s.width = width
	s.height = height
}

type SVGWriter interface {
	WriteSVG(io.Writer) error
}
