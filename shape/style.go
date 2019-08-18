package shape

import (
	"bytes"
	"io"
)

func NewStyler(dest io.Writer) *Styler {
	return &Styler{dest: dest}
}

type Styler struct {
	dest io.Writer

	err     error
	written int
}

func (styler *Styler) write(s []byte) {
	styler.written, styler.err = styler.dest.Write(s)
}

var styles = map[string]string{
	"highlight":      "stroke:red",
	"highlight-head": "stroke:red;fill:#ffffff",
	"arrow":          "stroke:black",
	"arrow-head":     "stroke:black;fill:#ffffff",
	"arrow-tail":     "fill:white;stroke:#d3d3d3",
	"line":           "stroke:#d3d3d3",
	"column-line":    "stroke:#d3d3d3",
	"record":         "fill:#ffffcc;stroke:black",
}

// Write adds a style attribute based on class.
// Limited to 1 class only
func (styler *Styler) Write(s []byte) (int, error) {
	field := []byte(`class="`)
	i := bytes.Index(s, field)
	if i == -1 {
		return styler.dest.Write(s)
	}
	write := styler.write
	write(s[:i])
	write(field)
	class := parseClass(s[i:])
	write(class)
	style, found := styles[string(class)]
	if found {
		write([]byte(`" style="`))
		write([]byte(style))
		write([]byte(`" `))
	}
	afterClass := i + len(field) + len(class) + 2
	write(s[afterClass:]) // the rest
	return styler.written, styler.err
}

// s should start with `class="NAME"...`
func parseClass(s []byte) []byte {
	j := len(`class="`)
	i := bytes.Index(s[j:], []byte(`"`))
	if i == -1 {
		return s[j:]
	}
	return s[j : j+i]
}
