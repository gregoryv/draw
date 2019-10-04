package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewLabel(text string) *Label {
	return &Label{
		Pos:   xy.Position{0, DefaultFont.LineHeight},
		Text:  text,
		Font:  DefaultFont,
		Pad:   DefaultPad,
		class: "label",
	}
}

type Label struct {
	Pos   xy.Position
	Text  string
	Font  Font
	Pad   Padding
	class string
}

func (l *Label) String() string {
	return fmt.Sprintf("label %s at %v", l.Text, l.Pos)
}

func (l *Label) Position() (int, int) { return l.Pos.XY() }
func (l *Label) SetX(x int)           { l.Pos.X = x }
func (l *Label) SetY(y int)           { l.Pos.Y = y }
func (l *Label) Width() int {
	return l.Font.TextWidth(l.Text)
}
func (l *Label) Height() int          { return l.Font.Height }
func (l *Label) Direction() Direction { return LR }
func (l *Label) SetClass(c string)    { l.class = c }

func (l *Label) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<text class="%s" x="%v" y="%v">%s</text>`,
		l.class, l.Pos.X, l.Pos.Y, l.Text)
	return err
}
