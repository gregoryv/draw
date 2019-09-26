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
		Class: "label",
	}
}

type Label struct {
	X, Y int
	Text string

	Font  Font
	Pad   Padding
	Class string
}

func (shape *Label) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<text class="%s" x="%v" y="%v">%s</text>`,
		shape.Class, shape.X, shape.Y, shape.Text)
	return err
}

func (l *Label) Height() int { return l.Font.Height }
func (l *Label) Width() int  { return len(l.Text) * l.Font.Width }

func (l *Label) SetX(x int)           { l.X = x }
func (l *Label) SetY(y int)           { l.Y = y }
func (l *Label) Direction() Direction { return LR }
func (l *Label) Position() (int, int) { return l.X, l.Y }
