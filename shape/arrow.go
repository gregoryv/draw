package shape

import (
	"fmt"
	"io"
	"math"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/xy"
)

func NewArrow(x1, y1, x2, y2 int) *Arrow {
	head := NewTriangle()
	head.SetX(x2)
	head.SetY(y2)
	head.SetClass("arrow-head")
	return &Arrow{
		Start: xy.Position{x1, y1},
		End:   xy.Position{x2, y2},
		Head:  head,
		class: "arrow",
	}
}

type Arrow struct {
	Start xy.Position
	End   xy.Position
	Tail  Shape
	Head  Shape
	class string
}

func (a *Arrow) String() string {
	return fmt.Sprintf("Arrow from %v to %v", a.Start, a.End)
}

func (a *Arrow) WriteSvg(out io.Writer) error {
	w, err := draw.NewTagWriter(out)
	x1, y1 := a.Start.XY()
	x2, y2 := a.End.XY()
	w.Printf(`<path class="%s" d="M%v,%v L%v,%v" />`, a.class, x1, y1, x2, y2)
	w.Print("\n")
	if a.Tail != nil {
		w.Printf(`<g transform="rotate(%v %v %v)">`, a.angle(), x1, y1)
		alignTail(a.Tail, x1, y1)
		a.Tail.SetClass(a.class + "-tail")
		a.Tail.WriteSvg(out)
		w.Print("</g>\n")
	}
	if a.Head != nil {
		w.Printf(`<g transform="rotate(%v %v %v)">`, a.angle(), x2, y2)
		a.Head.SetX(a.End.X)
		a.Head.SetY(a.End.Y)
		a.Head.SetClass(a.class + "-head")
		a.Head.WriteSvg(out)
		w.Print("</g>\n")
	}
	return *err
}

func alignTail(s Shape, x, y int) {
	switch s := s.(type) {
	case *Circle:
		s.SetX(x)
		s.SetY(y - s.Radius)
	default:
		s.SetX(x)
		s.SetY(y)
	}
}

func (arrow *Arrow) absAngle() float64 {
	return math.Abs(float64(arrow.angle()))
}

// angle returns degrees the head of an arrow should rotate depending
// on direction
func (a *Arrow) angle() int {
	var (
		start = a.Start
		end   = a.End
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
	case a.DirQ1():
		a := float64(end.Y - start.Y)
		b := float64(end.X - start.X)
		A := math.Atan(a / b)
		return radians2degrees(A)
	case a.DirQ2():
		a := float64(end.Y - start.Y)
		b := float64(start.X - end.X)
		A := math.Atan(a / b)
		return 180 - radians2degrees(A)
	case a.DirQ3():
		a := float64(start.Y - end.Y)
		b := float64(start.X - end.X)
		A := math.Atan(a / b)
		return radians2degrees(A) + 180
	case a.DirQ4():
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

func (a *Arrow) endpoints() (xy.Position, xy.Position) {
	return a.Start, a.End
}

func radians2degrees(A float64) int {
	return int(A * 180 / math.Pi)
}

func (a *Arrow) Height() int {
	return intAbs(a.Start.Y - a.End.Y)
}

func (a *Arrow) Width() int {
	return intAbs(a.Start.X - a.End.X)
}

func (a *Arrow) Position() (int, int) {
	return a.Start.XY()
}

func (a *Arrow) SetX(x int) {
	diff := a.Start.X - x
	a.Start.X = x
	a.End.X = a.End.X - diff // Set X2 so the entire arrow moves
}

func (a *Arrow) SetY(y int) {
	diff := a.Start.Y - y
	a.Start.Y = y
	a.End.Y = a.End.Y - diff // Set Y2 so the entire arrow moves
}

func (a *Arrow) Direction() Direction {
	if a.Start.LeftOf(a.End) {
		return LR
	}
	return RL
}

func (a *Arrow) SetClass(c string) { a.class = c }

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
