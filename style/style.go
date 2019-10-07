// Package style provides svg class base styles for design diagrams
package style

import (
	"bytes"
	"fmt"
	"io"
)

func NewStyler(dest io.Writer) *Style {
	return &Style{dest: dest}
}

type Style struct {
	dest    io.Writer
	err     error
	written int
	styles  map[string]string
}

func (style *Style) write(s []byte) {
	style.written, style.err = style.dest.Write(s)
}

// classname -> style
var DefaultStyle = map[string]string{
	"note":                  `font-family="Arial,Helvetica,sans-serif" font-size="12px"`,
	"note-box":              `stroke="black" fill="#ffffff"`,
	"highlight":             `stroke="red"`,
	"highlight-head":        `stroke="red" fill="#ffffff"`,
	"implements-arrow":      `stroke="black" stroke-dasharray="5,5,5"`,
	"implements-arrow-head": `stroke="black" fill="#ffffff"`,
	"arrow":                 `stroke="black"`,
	"arrow-head":            `stroke="black" fill="#ffffff"`,
	"arrow-tail":            `fill="white" stroke="#d3d3d3"`,
	"line":                  `stroke="#d3d3d3"`,
	"column-line":           `stroke="#d3d3d3"`,
	"record":                `stroke="black" fill="#ffffcc"`,
	"record-title":          `font-family="Arial" Helvetica="sans-serif" font-size="12px"`,
	"field":                 `font-family="Arial" Helvetica="sans-serif" font-size="12px"`,
	"method":                `font-family="Arial" Helvetica="sans-serif" font-size="12px"`,
	"record-label":          `font-family="Arial" Helvetica="sans-serif" font-size="12px"`,
	"label":                 `font-family="Arial" Helvetica="sans-serif" font-size="12px"`,
}

// Write adds a style attribute based on class. Limited to 1 class
// only and assumes the entire classname attribute is found.
func (style *Style) Write(p []byte) (int, error) {
	class, i := style.scanClass(p)
	if i == -1 {
		return style.dest.Write(p)
	}
	write := style.write
	s, found := style.styles[string(class)]
	if !found {
		s, found = DefaultStyle[string(class)]
	}
	if found {
		write([]byte(s))
	} else {
		write([]byte(`class="`))
		write(class)
		write([]byte(`"`))
	}
	write(p[i:]) // the rest
	return style.written, style.err
}

var field = []byte(`class="`)

// scanClass returns name of class and position after the attribute.
// position is -1 if no class was found. Everything up to the class,
// except the class attribute is written to the underlyinge writer.
func (style *Style) scanClass(p []byte) ([]byte, int) {
	i := bytes.Index(p, field)
	if i == -1 {
		return []byte{}, -1
	}
	style.write(p[:i])
	var (
		class = make([]byte, 0)
		j     int
		c     byte
		endOk bool
	)
	i = len(field) + i
	for j, c = range p[i:] {
		if c == '"' {
			endOk = true
			break
		}
		class = append(class, c)
	}
	if !endOk {
		panic(fmt.Sprintf("malformed: %s", string(p)))
	}
	j++
	return class, i + j
}
