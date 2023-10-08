package shape

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewLabel(text string) *Label {
	return &Label{
		text:  template.HTMLEscapeString(text),
		Font:  draw.DefaultFont,
		Pad:   draw.DefaultPad,
		class: "label",
	}
}

type Label struct {
	x int
	y int

	text string
	href string

	Font  draw.Font
	Pad   draw.Padding
	class string
}

func (l *Label) Text() string     { return l.text }
func (l *Label) SetText(v string) { l.text = v }

func (l *Label) String() string {
	return fmt.Sprintf("label %s at %v", l.text, &xy.Point{X: l.x, Y: l.y})
}

func (l *Label) SetHref(v string) { l.href = v }

func (l *Label) Position() (x int, y int) { return l.x, l.y }
func (l *Label) SetX(x int)               { l.x = x }
func (l *Label) SetY(y int)               { l.y = y }

func (l *Label) Width() int {
	return l.Font.TextWidth(longestLine(l.text))
}

func (l *Label) Height() int {
	return l.Font.LineHeight * (strings.Count(l.text, "\n") + 1)
}
func (l *Label) Direction() Direction { return DirectionRight }
func (l *Label) SetClass(c string)    { l.class = c }

func (l *Label) WriteSVG(out io.Writer) error {
	x, y := l.Position()
	y += l.Font.LineHeight
	w, err := nexus.NewPrinter(out)

	if l.href != "" {
		w.Printf(`<a href="%s">`, l.href)
	}
	// support multilines
	for i, line := range strings.Split(l.text, "\n") {
		if i > 0 { // write new line "after" each <text>, but not last
			w.Print("\n")
		}
		w.Printf(`<text class="%s" font-size="%vpx" x="%v" y="%v">%s</text>`,
			l.class, l.Font.Height, x, y+(l.Font.LineHeight*i), line)
	}
	if l.href != "" {
		w.Printf(`</a>`)
	}
	return *err
}

func (l *Label) Edge(start xy.Point) xy.Point {
	return boxEdge(start, l)
}

func longestLine(text string) string {
	lines := strings.Split(text, "\n")
	var max int
	var longest string
	for _, line := range lines {
		if len(line) < max {
			continue
		}
		max = len(line)
		longest = line
	}
	return longest
}
