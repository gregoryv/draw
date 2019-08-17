package shape

import (
	"io"
	"math"
)

type Arrow struct {
	X1, Y1 int
	X2, Y2 int

	Tail bool
}

func (arrow *Arrow) WriteSvg(w io.Writer) error {
	tag, _, err := newTagPrinter(w)
	x1, y1 := arrow.start()
	x2, y2 := arrow.end()
	tag.printf(`<path class="arrow" d="M%v,%v L%v,%v" />`, x1, y1, x2, y2)
	tag.printf("\n")
	if arrow.Tail {
		tag.printf(`<circle class="arrowtail" cx="%v" cy="%v" r="3" />`, x1, y1)
		tag.printf("\n")
	}
	tag.printf(`<g transform="rotate(%v %v %v)">`, arrow.angle(), x2, y2)
	tag.printf(`<path class="arrowhead" d="M%v,%v l-8,-4 l 0,8 Z" />`, x2, y2)
	tag.printf(`</g>`)
	tag.printf("\n")
	return *err
}

func (arrow *Arrow) angle() int {
	x1, y1 := arrow.start()
	x2, y2 := arrow.end()

	var (
		q1 = y1 < y2 && x1 < x2
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

func quadrantAngleDiffx(x1, y1, x2, y2 int) int {
	switch {
	case inQuadrant1(x1, y1, x2, y2):
		return 0
	case inQuadrant2(x1, y1, x2, y2):
		return 1
	case inQuadrant3(x1, y1, x2, y2):
		return 2
	case inQuadrant4(x1, y1, x2, y2):
		return 3
	}
	return 0
}

func inQuadrant1(x1, y1, x2, y2 int) bool { return y1 < y2 && x1 < x2 }
func inQuadrant2(x1, y1, x2, y2 int) bool { return y1 < y2 && x1 > x2 }
func inQuadrant3(x1, y1, x2, y2 int) bool { return y1 > y2 && x1 > x2 }
func inQuadrant4(x1, y1, x2, y2 int) bool { return y1 > y2 && x1 < x2 }

func (arrow *Arrow) start() (int, int) { return arrow.X1, arrow.Y1 }
func (arrow *Arrow) end() (int, int)   { return arrow.X2, arrow.Y2 }

func (arrow *Arrow) Height() int {
	return intAbs(arrow.Y1 - arrow.Y2)
}

func (arrow *Arrow) Width() int {
	return intAbs(arrow.X1 - arrow.X2)
}

func (arrow *Arrow) Position() (int, int) {
	return arrow.X1, arrow.Y1
}

func (arrow *Arrow) SetX(x int) {
	diff := arrow.X1 - x
	arrow.X1 = x
	arrow.X2 = arrow.X2 - diff // Set X2 so the entire arrow moves
}

func (arrow *Arrow) SetY(y int) {
	diff := arrow.Y1 - y
	arrow.Y1 = y
	arrow.Y2 = arrow.Y2 - diff // Set Y2 so the entire arrow moves
}

func (arrow *Arrow) Direction() Direction {
	if arrow.X1 <= arrow.X2 {
		return LR
	}
	return RL
}
