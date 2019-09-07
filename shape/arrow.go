package shape

import (
	"io"
	"math"

	"github.com/gregoryv/go-design/xy"
)

func NewArrow(x1, y1, x2, y2 int) *Arrow {
	return &Arrow{
		Start: xy.Position{x1, y1},
		End:   xy.Position{x2, y2},
	}
}

type Arrow struct {
	Start xy.Position
	End   xy.Position

	Tail  bool
	Class string
}

func (arrow *Arrow) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	x1, y1 := arrow.Start.XY()
	x2, y2 := arrow.End.XY()
	w.printf(`<path class="%s" d="M%v,%v L%v,%v" />`, arrow.class(), x1, y1, x2, y2)
	w.print("\n")
	if arrow.Tail {
		w.printf(`<circle class="%s-tail" cx="%v" cy="%v" r="3" />`, arrow.class(), x1, y1)
		w.print("\n")
	}
	w.printf(`<g transform="rotate(%v %v %v)">`, arrow.angle(), x2, y2)
	// the path is drawn as if it points straight to the right
	w.printf(`<path class="%s-head" d="M%v,%v l-8,-4 l 0,8 Z" />`, arrow.class(), x2, y2)
	w.print("</g>\n")
	return *err
}

func (arrow *Arrow) angle() int {
	start := arrow.Start
	end := arrow.End
	x1, y1 := arrow.Start.XY()
	x2, y2 := arrow.End.XY()

	var (
		// quadrandts start at top right and are counted counter clockwise
		q1 = start.LeftOf(end) && start.Above(end)
		q2 = y1 < y2 && x1 > x2
		q3 = y1 > y2 && x1 > x2
		q4 = y1 > y2 && x1 < x2
		// straight arrows
		right = x1 < x2 && y1 == y2
		left  = x1 > x2 && y1 == y2
		down  = x1 == x2 && y1 < y2
		up    = x1 == x2 && y1 > y2
	)

	switch {
	case right: // most frequent arrow on top
	case left:
		return 180
	case down:
		return 90
	case up:
		return -90
	case q1:
		a := float64(y2 - y1)
		b := float64(x2 - x1)
		A := math.Asin(math.Abs(a) / math.Abs(b))
		return int(A * 180 / math.Pi)
	case q2:
		a := float64(y2 - y1)
		b := float64(x1 - x2)
		A := math.Asin(math.Abs(b) / math.Abs(a))
		return int(A*180/math.Pi) + 90
	case q3:
		a := float64(y1 - y2)
		b := float64(x1 - x2)
		A := math.Asin(math.Abs(a) / math.Abs(b))
		return int(A*180/math.Pi) + 180
	case q4:
		a := float64(y1 - y2)
		b := float64(x2 - x1)
		A := math.Asin(math.Abs(a) / math.Abs(b))
		return -int(A * 180 / math.Pi)
	}

	return 0
}

func (arrow *Arrow) Height() int {
	return intAbs(arrow.Start.Y - arrow.End.Y)
}

func (arrow *Arrow) Width() int {
	return intAbs(arrow.Start.X - arrow.End.X)
}

func (arrow *Arrow) Position() (int, int) {
	return arrow.Start.XY()
}

func (arrow *Arrow) SetX(x int) {
	diff := arrow.Start.X - x
	arrow.Start.X = x
	arrow.End.X = arrow.End.X - diff // Set X2 so the entire arrow moves
}

func (arrow *Arrow) SetY(y int) {
	diff := arrow.Start.Y - y
	arrow.Start.Y = y
	arrow.End.Y = arrow.End.Y - diff // Set Y2 so the entire arrow moves
}

func (arrow *Arrow) Direction() Direction {
	if arrow.Start.LeftOf(arrow.End) {
		return LR
	}
	return RL
}
func (arrow *Arrow) class() string {
	if arrow.Class == "" {
		return "arrow"
	}
	return arrow.Class
}
