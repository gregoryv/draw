package shape

import (
	"bytes"
	"io"
)

type Styler struct {
	dest io.Writer

	err     error
	written int
}

func (styler *Styler) write(s []byte) {
	styler.written, styler.err = styler.dest.Write(s)
}

var styles = map[string]string{
	"arrowhead": "stroke:green;fill:green",
	"arrow":     "stroke:black",
	"line":      "stroke:black",
	"record":    "fill:#ffffcc;stroke:black",
	"arrowtail": "fill:red;stroke:black",
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
