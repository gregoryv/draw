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

// inline me returns SVG with inlined style.
func inline(w draw.SVGWriter, style draw.Style) string {
	var buf bytes.Buffer
	style.SetOutput(&buf)
	w.WriteSVG(&style)
	return buf.String()
}
