package shape

import (
	"fmt"
	"io"
)

func NewLabel(text string) *Label {
	return &Label{
		X: 0, Y: 0, Text: text,
		Font:  DefaultFont,
		Pad:   DefaultPad,
		class: "label",
	}
}

type Label struct {
	X, Y int
	Text string

	Font  Font
	Pad   Padding
	class string
}

func (l *Label) String() string {
	return fmt.Sprintf("label %s at %v,%v", l.Text, l.X, l.Y)
}

func (l *Label) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<text class="%s" x="%v" y="%v">%s</text>`,
		l.class, l.X, l.Y, l.Text)
	return err
}

func (l *Label) Height() int { return l.Font.Height }
func (l *Label) Width() int  { return l.Font.TextWidth(l.Text) }

func (l *Label) SetX(x int)           { l.X = x }
func (l *Label) SetY(y int)           { l.Y = y }
func (l *Label) Direction() Direction { return LR }
func (l *Label) Position() (int, int) { return l.X, l.Y }

func (l *Label) SetClass(c string) { l.class = c }
