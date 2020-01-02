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

func (l *Line) String() string {
	return fmt.Sprintf("Line from %v to %v", l.Start, l.End)
}

func (l *Line) WriteSvg(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		`<line class="%s" x1="%v" y1="%v" x2="%v" y2="%v"/>`,
		l.class,
		l.Start.X, l.Start.Y,
		l.End.X, l.End.Y,
	)
	return err
}

func (l *Line) Position() (int, int) {
	return l.Start.XY()
}

func (l *Line) Width() int {
	return intAbs(l.Start.X - l.End.X)
}

func (l *Line) Height() int {
	return intAbs(l.Start.Y - l.End.Y)
}

func (l *Line) SetX(x int) {
	diff := l.Start.X - x
	l.Start.X = x
	l.End.X = l.End.X - diff
}

func (l *Line) SetY(y int) {
	diff := l.Start.Y - y
	l.Start.Y = y
	l.End.Y = l.End.Y - diff
}

func (l *Line) Direction() Direction {
	if l.Start.X <= l.End.X {
		return RightDir
	}
	return LeftDir
}

func (l *Line) SetClass(c string) { l.class = c }
