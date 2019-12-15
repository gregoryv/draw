// Package draw provides svg diagram creators
package design

import (
	"io"
	"os"

	"github.com/gregoryv/draw/shape"
)

type SvgWriter interface {
	WriteSvg(io.Writer) error
}

// saveAs saves diagram using default style to filename
func saveAs(dia SvgWriter, style shape.Style, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	style.SetOutput(fh)
	return dia.WriteSvg(&style)
}
