package draw

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// NewStyle returns a style based on the default values,
// eg. draw.DefaultFont in this package.
func NewStyle() Style {
	return Style{
		Font:    DefaultFont,
		TextPad: DefaultTextPad,
		Pad:     DefaultPad,
		Spacing: DefaultSpacing,
		dest:    ioutil.Discard,
	}
}

type Style struct {
	Font
	TextPad Padding // Surrounding text
	Pad     Padding // E.g. records
	Spacing int     // Between shapes in e.g. diagrams
	dest    io.Writer
	err     error
	written int
	styles  map[string]string
}

var (
	DefaultFont    = Font{Height: 12, LineHeight: 16, charWidths: arial}
	DefaultTextPad = Padding{Left: 6, Top: 4, Bottom: 6, Right: 10}
	DefaultPad     = Padding{Left: 10, Top: 2, Bottom: 7, Right: 10}
	DefaultSpacing = 30 // between elements
)

// ClassAttributes define mapping between classes and svg attributes.
// Setting attributes that modify size or position is not advised.
type ClassAttributes map[string]string

// CSS returns cascading rules for embedding in html
func (me ClassAttributes) CSS() string {
	var s strings.Builder
	for class, rules := range me {
		s.WriteString(".")
		s.WriteString(class)
		s.WriteString(" {\n")
		rules = strings.TrimSpace(rules) + " "
		r := strings.Split(rules, `" `)
		for _, rule := range r[:len(r)-1] {
			s.WriteString("  ")
			s.WriteString(rule)
			s.WriteString("\";\n")
		}
		s.WriteString("}\n\n")
	}
	return s.String()
}

var DefaultClassAttributes = ClassAttributes{
	"area-red-label":   `font-style="italic" font-family="Arial,Helvetica,sans-serif"`,
	"area-green-label": `font-style="italic" font-family="Arial,Helvetica,sans-serif"`,
	"area-blue-label":  `font-style="italic" font-family="Arial,Helvetica,sans-serif"`,
	"area-red":         `stroke="black" stroke-width="0" fill="#ff9999" fill-opacity="0.1"`,
	"area-green":       `stroke="black" stroke-width="0" fill="#ccff99" fill-opacity="0.1"`,
	"area-blue":        `stroke="black" stroke-width="0" fill="#99e6ff" fill-opacity="0.1"`,

	"actor":                 `stroke="black" stroke-width="2" fill="#ffffff"`,
	"circle":                `stroke="#d3d3d3" stroke-width="2" fill="#ffffff"`,
	"cylinder":              `stroke="#d3d3d3" stroke-width="1" fill="#ffffff"`,
	"card":                  `stroke="#d3d3d3" stroke-width="1" fill="#ffffff"`,
	"card-external":         `stroke="#d3d3d3" stroke-width="1" fill="#f8f9fa"`,	
	"card-title":            `font-family="Arial,Helvetica,sans-serif" font-weight="bold"`,
	"database":              `stroke="#d3d3d3" stroke-width="1" fill="#ffffff"`,
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
	"dashed-arrow":          `stroke="black"`,
	"dashed-arrow-head":     `stroke="black" fill="#ffffff"`,
	"dashed-arrow-tail":     `stroke="black" fill="#777777"`,
	"activity-arrow":        `stroke="black"`,
	"activity-arrow-head":   `stroke="black" fill="#ffffff"`,
	"activity-arrow-tail":   `stroke="black" fill="#777777"`,
	"compose-arrow":         `stroke="black"`,
	"compose-arrow-head":    `stroke="black" fill="#ffffff"`,
	"compose-arrow-tail":    `stroke="black" fill="#777777"`,
	"aggregate-arrow":       `stroke="black"`,
	"aggregate-arrow-head":  `stroke="black" fill="#ffffff"`,
	"aggregate-arrow-tail":  `stroke="black" fill="#ffffff"`,
	"external":              `stroke="#d3d3d3" fill="#e2e2e2"`,
	"dim":                   `stroke="#d3d3d3" fill="#e2e2e2"`,
	"hexagon":               `stroke="#d3d3d3" fill="#ffffff"`,
	"hexagon-title":         `font-family="Arial,Helvetica,sans-serif"`,
	"internet":              `stroke="#d3d3d3" fill="#e2e2e2"`,
	"internet-title":        `font-family="Arial,Helvetica,sans-serif"`,
	"line":                  `stroke="black"`,
	"triangle":              `stroke="black"`,
	"column-line":           `stroke="#d3d3d3"`,
	"process":               `stroke="#d3d3d3" fill="#ffffff"`,
	"process-title":         `font-family="Arial,Helvetica,sans-serif"`,
	"record":                `stroke="#d3d3d3" fill="#ffffff"`,
	"record-line":           `stroke="#d3d3d3"`,
	"record-title":          `font-family="Arial,Helvetica,sans-serif"`,
	"rect":                  `stroke="#d3d3d3" fill="#ffffff"`,
	"rect-title":            `font-family="Arial,Helvetica,sans-serif"`,
	"root":                  `font-family="Arial,Helvetica,sans-serif"`, // root svg tag
	"skip":                  `stroke="#ffffff" stroke-dasharray="2,2,2"`,
	"span-green":            `stroke="#d3d3d3" fill="#ccff99" rx="5" ry="5"`,
	"span-green-title":      `font-family="Arial,Helvetica,sans-serif"`,
	"span-blue":             `stroke="#d3d3d3" fill="#99e6ff" rx="5" ry="5"`,
	"span-blue-title":       `font-family="Arial,Helvetica,sans-serif"`,
	"span-red":              `stroke="#d3d3d3" fill="#ff9999" rx="5" ry="5"`,
	"span-red-title":        `font-family="Arial,Helvetica,sans-serif"`,
	"state-title":           `font-family="Arial,Helvetica,sans-serif"`,
	"state":                 `stroke="#d3d3d3" fill="#ffffff" rx="10" ry="10"`,
	"store":                 `stroke="#d3d3d3" fill="#ffffff"`,
	"store-title":           `font-family="Arial,Helvetica,sans-serif"`,
	"component":             `stroke="#d3d3d3" fill="#ffffff"`,
	"component-title":       `font-family="Arial,Helvetica,sans-serif"`,
	"field":                 `font-family="Arial,Helvetica,sans-serif"`,
	"method":                `font-family="Arial,Helvetica,sans-serif"`,
	"record-label":          `font-family="Arial,Helvetica,sans-serif"`,
	"label":                 `font-family="Arial,Helvetica,sans-serif"`,
	"weekend":               `font-family="Arial,Helvetica,sans-serif" fill="#f3f3f3"`,
	"weekend-title":         `font-family="Arial,Helvetica,sans-serif"`,
	"caption":               `font-family="Arial,Helvetica,sans-serif"`,
	"diamond":               `stroke="#d3d3d3" fill="#333333"`,
	"decision":              `stroke="#d3d3d3" fill="#ffffff"`,
}

// Write adds a style attribute based on class. Limited to 1 class
// only and assumes the entire classname attribute is found.
func (s *Style) Write(p []byte) (int, error) {
	s.written = 0
	class, i := s.scanClass(p)
	if i == -1 {
		return s.dest.Write(p)
	}
	write := s.write
	st, found := s.styles[string(class)]
	if !found {
		st, found = DefaultClassAttributes[string(class)]
	}
	if found {
		write([]byte(st))
	} else {
		write([]byte(`class="`))
		write(class)
		write([]byte(`"`))
	}
	write(p[i:]) // the rest
	return s.written, s.err
}

func (s *Style) write(b []byte) {
	if s.err != nil {
		return
	}
	n, err := s.dest.Write(b)
	s.written += n
	s.err = err
}

var field = []byte(`class="`)

// scanClass returns name of class and position after the attribute.
// position is -1 if no class was found. Everything up to the class,
// except the class attribute is written to the underlyinge writer.
func (s *Style) scanClass(p []byte) ([]byte, int) {
	i := bytes.Index(p, field)
	if i == -1 {
		return []byte{}, -1
	}
	s.write(p[:i])
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

// SetOutput sets the destination of calls to Write.
func (s *Style) SetOutput(w io.Writer) {
	if w == nil {
		w = ioutil.Discard
	}
	s.dest = w
}
