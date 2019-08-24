package shape

import (
	"fmt"
	"io"
)

type Label struct {
	X, Y int
	Text string

	Font Font
	Pad  Padding
}

func (shape *Label) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<text x="%v" y="%v">%s</text>`,
		shape.X, shape.Y, shape.Text)
	return err
}

func (l *Label) Height() int { return l.Font.Height }
func (l *Label) Width() int  { return len(l.Text) * l.Font.Width }

func (l *Label) SetX(x int)           { l.X = x }
func (l *Label) SetY(y int)           { l.Y = y }
func (l *Label) Direction() Direction { return LR }
func (l *Label) Position() (int, int) { return l.X, l.Y }
