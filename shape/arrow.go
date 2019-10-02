package shape

import (
	"fmt"
	"io"
	"math"

	"github.com/gregoryv/go-design/xy"
)

func NewArrow(x1, y1, x2, y2 int) *Arrow {
	return &Arrow{
		Start: xy.Position{x1, y1},
		End:   xy.Position{x2, y2},
		Head:  NewTriangle(x2, y2, "arrow-head"),
	}
}

type Arrow struct {
	Start xy.Position
	End   xy.Position

	Tail  bool
	Head  Shape
	Class string
}

func (a *Arrow) String() string {
	return fmt.Sprintf("Arrow from %v to %v", a.Start, a.End)
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
	// Update position before rendering
	arrow.Head.SetX(arrow.End.X)
	arrow.Head.SetY(arrow.End.Y)
	arrow.Head.WriteSvg(out)
	w.print("</g>\n")
	return *err
}

func (arrow *Arrow) absAngle() float64 {
	return math.Abs(float64(arrow.angle()))
}

// angle returns degrees the head of an arrow should rotate depending
// on direction
func (arrow *Arrow) angle() int {
	var (
		start = arrow.Start
		end   = arrow.End
		// straight arrows
		right = start.LeftOf(end) && start.Y == end.Y
		left  = start.RightOf(end) && start.Y == end.Y
		down  = start.Above(end) && start.X == end.X
		up    = start.Below(end) && start.X == end.X
	)
	switch {
	case right: // most frequent arrow on top
	case left:
		return 180
	case down:
		return 90
	case up:
		return -90
	case arrow.DirQ1():
		a := float64(end.Y - start.Y)
		b := float64(end.X - start.X)
		A := math.Atan(a / b)
		return radians2degrees(A)
	case arrow.DirQ2():
		a := float64(end.Y - start.Y)
		b := float64(start.X - end.X)
		A := math.Atan(a / b)
		return 180 - radians2degrees(A)
	case arrow.DirQ3():
		a := float64(start.Y - end.Y)
		b := float64(start.X - end.X)
		A := math.Atan(a / b)
		return radians2degrees(A) + 180
	case arrow.DirQ4():
		a := float64(start.Y - end.Y)
		b := float64(end.X - start.X)
		A := math.Atan(a / b)
		return -radians2degrees(A)
	}
	return 0
}

// DirQ1 returns true if the arrow points to the bottom-right
// quadrant.
func (a *Arrow) DirQ1() bool {
	start, end := a.endpoints()
	return start.LeftOf(end) && end.Below(start)
}

// DirQ2 returns true if the arrow points to the bottom-left
// quadrant.
func (a *Arrow) DirQ2() bool {
	start, end := a.endpoints()
	return start.RightOf(end) && end.Below(start)
}

// DirQ3 returns true if the arrow points to the top-left
// quadrant.
func (a *Arrow) DirQ3() bool {
	start, end := a.endpoints()
	return start.RightOf(end) && end.Above(start)
}

// DirQ4 returns true if the arrow points to the top-right
// quadrant.
func (a *Arrow) DirQ4() bool {
	start, end := a.endpoints()
	return start.LeftOf(end) && end.Above(start)
}

func (arrow *Arrow) endpoints() (xy.Position, xy.Position) {
	return arrow.Start, arrow.End
}

func radians2degrees(A float64) int {
	return int(A * 180 / math.Pi)
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

func NewArrowBetween(a, b Shape) *Arrow {
	ax, ay := a.Position()
	bx, by := b.Position()

	// From center to center
	x1 := ax + a.Width()/2
	y1 := ay + a.Height()/2
	x2 := bx + b.Width()/2
	y2 := by + b.Height()/2
	arrow := NewArrow(x1, y1, x2, y2)
	bs, ok := b.(Edge)
	if ok {
		p := bs.Edge(arrow.Start)
		arrow.End.X = p.X
		arrow.End.Y = p.Y
	}
	as, ok := a.(Edge)
	if ok {
		p := as.Edge(arrow.End)
		arrow.Start.X = p.X
		arrow.Start.Y = p.Y
	}

	return arrow
}
