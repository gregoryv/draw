/*
Package draw provides svg writing features.

	s := NewSvg()
	s.WriteSvg(os.Stdout)

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

func (s *Svg) WriteSvg(out io.Writer) error {
	w, err := NewTagWriter(out)
	w.Printf(`<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  width="%v" height="%v" font-family="Arial, Helvetica, sans-serif">`, s.width, s.height)

	for _, c := range s.Content {
		w.Print("\n")
		c.WriteSvg(w)
	}
	w.Print("</svg>")
	return *err
}

func (s *Svg) Append(w ...SvgWriter) {
	s.Content = append(s.Content, w...)
}

func (s *Svg) Prepend(w ...SvgWriter) {
	s.Content = append(w, s.Content...)
}

func (s *Svg) Width() int  { return s.width }
func (s *Svg) Height() int { return s.height }

func (s *Svg) SetWidth(w int)   { s.width = w }
func (s *Svg) SetHeight(h int)  { s.height = h }
func (s *Svg) SetSize(w, h int) { s.SetWidth(w); s.SetHeight(h) }

type SvgWriter interface {
	WriteSvg(io.Writer) error
}
