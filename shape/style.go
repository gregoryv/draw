package shape

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func NewStyle(dest io.Writer) Style {
	return Style{
		Font:    DefaultFont,
		TextPad: DefaultTextPad,
		Pad:     DefaultPad,
		dest:    dest,
	}
}

type Style struct {
	Font
	TextPad Padding // Surrounding text
	Pad     Padding // E.g. records
	dest    io.Writer
	err     error
	written int
	styles  map[string]string
}

var (
	DefaultFont    = Font{Height: 12, LineHeight: 16, charWidths: arial}
	DefaultTextPad = Padding{Left: 6, Top: 4, Bottom: 6, Right: 10}
	DefaultPad     = Padding{Left: 10, Top: 2, Bottom: 7, Right: 10}
)

// ClassAttributes define mapping between classes and svg attributes.
// Setting attributes that modify size or position is not advised.
var ClassAttributes = map[string]string{
	"circle":                `stroke="black" stroke-width="2" fill="#ffffff"`,
	"dot":                   `stroke="black"`,
	"exit":                  `stroke="black" stroke-width="2" fill="#ffffff"`,
	"exit-dot":              `stroke="black"`,
	"note":                  `font-family="Arial,Helvetica,sans-serif"`,
	"note-box":              `stroke="#d3d3d3" fill="#ffffcc"`,
	"highlight":             `stroke="red"`,
	"highlight-head":        `stroke="red" fill="#ffffff"`,
	"implements-arrow":      `stroke="black" stroke-dasharray="5,5,5"`,
	"implements-arrow-head": `stroke="black" fill="#ffffff"`,
	"arrow":                 `stroke="black"`,
	"arrow-head":            `stroke="black" fill="#ffffff"`,
	"arrow-tail":            `stroke="black" fill="#777777"`,
	"compose-arrow":         `stroke="black"`,
	"compose-arrow-head":    `stroke="black" fill="#ffffff"`,
	"compose-arrow-tail":    `stroke="black" fill="#777777"`,
	"line":                  `stroke="black"`,
	"column-line":           `stroke="#d3d3d3"`,
	"record":                `stroke="#d3d3d3" fill="#ffffff"`,
	"record-line":           `stroke="#d3d3d3"`,
	"record-title":          `font-family="Arial,Helvetica,sans-serif"`,
	"rect":                  `stroke="#d3d3d3" fill="#ffffff"`,
	"state-title":           `font-family="Arial,Helvetica,sans-serif"`,
	"state":                 `stroke="#d3d3d3" fill="#ffffff" rx="10" ry="10"`,
	"component":             `stroke="#d3d3d3" fill="#ffffff"`,
	"component-title":       `font-family="Arial,Helvetica,sans-serif"`,
	"field":                 `font-family="Arial,Helvetica,sans-serif"`,
	"method":                `font-family="Arial,Helvetica,sans-serif"`,
	"record-label":          `font-family="Arial,Helvetica,sans-serif"`,
	"label":                 `font-family="Arial,Helvetica,sans-serif"`,
	"caption":               `font-family="Arial,Helvetica,sans-serif"`,
	"diamond":               `stroke="#d3d3d3" fill="#333333"`,
}

// Write adds a style attribute based on class. Limited to 1 class
// only and assumes the entire classname attribute is found.
func (style *Style) Write(p []byte) (int, error) {
	style.written = 0
	class, i := style.scanClass(p)
	if i == -1 {
		return style.dest.Write(p)
	}
	write := style.write
	s, found := style.styles[string(class)]
	if !found {
		s, found = ClassAttributes[string(class)]
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

func (style *Style) write(s []byte) {
	if style.err != nil {
		return
	}
	n, err := style.dest.Write(s)
	style.written += n
	style.err = err
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

func (s *Style) SetOutput(out io.Writer) {
	if out == nil {
		out = ioutil.Discard
	}
	s.dest = out
}
