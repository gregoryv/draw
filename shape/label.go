package shape

import (
	"fmt"
	"html/template"
	"io"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewLabel(text string) *Label {
	return &Label{
		Text:  template.HTMLEscapeString(text),
		Font:  DefaultFont,
		Pad:   DefaultPad,
		class: "label",
	}
}

type Label struct {
	x int
	y int

	Text string
	href string

	Font  Font
	Pad   Padding
	class string
}

func (l *Label) String() string {
	return fmt.Sprintf("label %s at %v", l.Text, &xy.Point{X: l.x, Y: l.y})
}

func (l *Label) SetHref(v string) { l.href = v }

func (l *Label) Position() (int, int) {
	return l.x, l.y
}

func (l *Label) SetX(x int) { l.x = x }
func (l *Label) SetY(y int) { l.y = y }
func (l *Label) Width() int {
	return l.Font.TextWidth(l.Text)
}

func (l *Label) Height() int          { return l.Font.LineHeight }
func (l *Label) Direction() Direction { return DirectionRight }
func (l *Label) SetClass(c string)    { l.class = c }

func (l *Label) WriteSVG(out io.Writer) error {
	x, y := l.Position()
	y += l.Font.LineHeight
	w, err := nexus.NewPrinter(out)

	if l.href != "" {
		w.Printf(`<a href="%s">`, l.href)
	}
	w.Printf(`<text class="%s" font-size="%vpx" x="%v" y="%v">%s</text>`,
		l.class, l.Font.Height, x, y, l.Text)
	if l.href != "" {
		w.Printf(`</a>`)
	}
	return *err
}

func (l *Label) Edge(start xy.Point) xy.Point {
	return boxEdge(start, l)
}
