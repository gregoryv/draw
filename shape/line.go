package shape

import (
	"fmt"
	"io"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

func (line *Line) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<line class="line" x1="%v" y1="%v" x2="%v" y2="%v"/>`,
		line.X1, line.Y1, line.X2, line.Y2,
	)
	return err
}

func (line *Line) Height() int {
	return intAbs(line.Y1 - line.Y2)
}

func (line *Line) Width() int {
	return intAbs(line.X1 - line.X2)
}

func (line *Line) Position() (int, int) {
	return line.X1, line.Y1
}

func (line *Line) SetX(x int) {
	diff := line.X1 - x
	line.X1 = x
	line.X2 = line.X2 - diff // Set X2 so the entire arrow moves
}

func (line *Line) SetY(y int) {
	diff := line.Y1 - y
	line.Y1 = y
	line.Y2 = line.Y2 - diff // Set Y2 so the entire arrow moves
}
