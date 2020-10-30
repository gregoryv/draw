// Package design provides diagram creators
package design

import (
	"bytes"
	"os"

	"github.com/gregoryv/draw"
)

// saveAs saves diagram with inlined style to the given filename
func saveAs(dia draw.SVGWriter, style draw.Style, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	style.SetOutput(fh)
	return dia.WriteSVG(&style)
}

func toString(d draw.SVGWriter) string {
	var buf bytes.Buffer
	d.WriteSVG(&buf)
	return buf.String()
}
