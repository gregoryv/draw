package shape

import (
	"fmt"
	"io"
)

type Label struct {
	X, Y int
	Text string
}

func (shape *Label) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<text x="%v" y="%v">%s</text>`,
		shape.X, shape.Y, shape.Text)
	return err
}

func (shape *Label) Height() int {
	fontHeight := 10
	return fontHeight
}

func (label *Label) Width() int {
	fontWidth := 10
	return len(label.Text) * fontWidth
}

func (shape *Label) SetX(x int) { shape.X = x }
func (shape *Label) SetY(y int) { shape.Y = y }

func (shape *Label) Position() (int, int) {
	return shape.X, shape.Y
}
