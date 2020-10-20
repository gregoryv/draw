package draw

import "bytes"

// Inline me returns SVG with inlined style.
func Inline(w SVGWriter, style Style) string {
	var buf bytes.Buffer
	style.SetOutput(&buf)
	w.WriteSVG(&style)
	return buf.String()
}
