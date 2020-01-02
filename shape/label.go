package shape

import (
	"fmt"
	"html/template"
	"io"

	"github.com/gregoryv/draw/xy"
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

	Text  string
	Font  Font
	Pad   Padding
	class string
}

func (l *Label) String() string {
	return fmt.Sprintf("label %s at %v", l.Text, &xy.Position{X: l.x, Y: l.y})
}

func (l *Label) Position() (int, int) {
	return l.x, l.y
}

func (l *Label) SetX(x int) { l.x = x }
func (l *Label) SetY(y int) { l.y = y }
func (l *Label) Width() int {
	return l.Font.TextWidth(l.Text)
}

func (l *Label) Height() int          { return l.Font.LineHeight }
func (l *Label) Direction() Direction { return RightDir }
func (l *Label) SetClass(c string)    { l.class = c }

func (l *Label) WriteSvg(w io.Writer) error {
	x, y := l.Position()
	y += l.Font.LineHeight
	_, err := fmt.Fprintf(w,
		`<text class="%s" font-size="%vpx" x="%v" y="%v">%s</text>`,
		l.class, l.Font.Height, x, y, l.Text)
	return err
}

func (l *Label) Edge(start xy.Position) xy.Position {
	return boxEdge(start, l)
}
