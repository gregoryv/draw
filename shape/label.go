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

func (label *Label) Height() int {
	return label.Font.Height
}

func (label *Label) Width() int {
	return len(label.Text) * label.Font.Width
}

func (label *Label) SetX(x int) { label.X = x }
func (label *Label) SetY(y int) { label.Y = y }

func (label *Label) Position() (int, int) {
	return label.X, label.Y
}

func (label *Label) Direction() Direction { return LR }
