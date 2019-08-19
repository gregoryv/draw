package design

import (
	"io"
	"os"

	"github.com/gregoryv/go-design/shape"
)

// WriteDefault applies default style while writing diagram to w
func WriteDefault(dia shape.SvgWriter, w io.Writer) error {
	return dia.WriteSvg(NewStyler(w))
}

// saveAs saves diagram using default style to filename
func saveAs(dia shape.SvgWriter, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	return dia.WriteSvg(NewStyler(fh))
}
