// Package design provides svg diagram creators
package design

import (
	"io"
	"os"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/shape"
)

type svgWriter interface {
	WriteSVG(io.Writer) error
}

// saveAs saves diagram using default style to filename
func saveAs(dia svgWriter, style shape.Style, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	style.SetOutput(fh)
	return dia.WriteSVG(&style)
}

func inlineSVG(w io.Writer, d draw.SVGWriter, s *shape.Style) error {
	s.SetOutput(w)
	return d.WriteSVG(s)
}
