package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/draw/xy"
)

func NewLine(x1, y1 int, x2, y2 int) *Line {
	return &Line{
		Start: xy.Position{x1, y1},
		End:   xy.Position{x2, y2},
		class: "line",
	}
}

type Line struct {
	Start xy.Position
	End   xy.Position

	class string
}

func (line *Line) String() string {
	return fmt.Sprintf("Line from %v to %v", line.Start, line.End)
}

func (line *Line) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<line class="%s" x1="%v" y1="%v" x2="%v" y2="%v"/>`,
		line.class,
		line.Start.X, line.Start.Y,
		line.End.X, line.End.Y,
	)
	return err
}

func (line *Line) Position() (int, int) {
	return line.Start.XY()
}

func (line *Line) Width() int {
	return intAbs(line.Start.X - line.End.X)
}

func (line *Line) Height() int {
	return intAbs(line.Start.Y - line.End.Y)
}

func (line *Line) SetX(x int) {
	diff := line.Start.X - x
	line.Start.X = x
	line.End.X = line.End.X - diff
}

func (line *Line) SetY(y int) {
	diff := line.Start.Y - y
	line.Start.Y = y
	line.End.Y = line.End.Y - diff
}

func (line *Line) Direction() Direction {
	if line.Start.X <= line.End.X {
		return LR
	}
	return RL
}

func (line *Line) SetClass(c string) { line.class = c }
