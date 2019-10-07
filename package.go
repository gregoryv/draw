// Package go-design provides svg diagram creators
package design

import (
	"io"
	"os"

	"github.com/gregoryv/go-design/shape"
)

type SvgWriter interface {
	WriteSvg(io.Writer) error
}

// saveAs saves diagram using default style to filename
func saveAs(dia SvgWriter, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	return dia.WriteSvg(shape.NewStyle(fh))
}
