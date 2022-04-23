// Package draw provides SVG writing features.
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

// WriteSVG writes <svg> </svg> tags and it's content to the given
// writer.
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

// SetWidth sets the SVG width in pixels.
func (s *SVG) SetWidth(w int) { s.width = w }

// SetHeight sets the SVG height in pixels.
func (s *SVG) SetHeight(h int) { s.height = h }

// SetSize sets the SVG size in pixels
func (s *SVG) SetSize(width, height int) {
	s.width = width
	s.height = height
}

type SVGWriter interface {
	WriteSVG(io.Writer) error
}
