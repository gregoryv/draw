package shape

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/gregoryv/draw/xy"
	"github.com/gregoryv/nexus"
)

func NewArrow(x1, y1, x2, y2 int) *Line {
	head := NewTriangle()
	head.SetX(x2)
	head.SetY(y2)
	head.SetClass("arrow-head")
	return &Line{
		Start: xy.Point{x1, y1},
		End:   xy.Point{x2, y2},
		Head:  head,
		class: "arrow",
	}
}

func NewLine(x1, y1 int, x2, y2 int) *Line {
	return &Line{
		Start: xy.Point{x1, y1},
		End:   xy.Point{x2, y2},
		class: "line",
	}
}

type Line struct {
	Start xy.Point
	End   xy.Point
	Tail  Shape
	Head  Shape

	class string
}

func (a *Line) String() string {
	return fmt.Sprintf("Arrow from %v to %v", a.Start, a.End)
}

func (a *Line) WriteSVG(out io.Writer) error {
	w, err := nexus.NewPrinter(out)
	x1, y1 := a.Start.XY()
	x2, y2 := a.End.XY()
	var dashed string
	if strings.Contains(a.class, "dashed") {
		dashed = `stroke-dasharray="5,5" `
	}
	w.Printf(`<path %sclass="%s" d="M%v,%v L%v,%v" />`, dashed, a.class, x1, y1, x2, y2)
	w.Print("\n")
	if a.Tail != nil {
		w.Printf(`<g transform="rotate(%v %v %v)">`, a.angle(), x1, y1)
		alignTail(a.Tail, x1, y1)
		a.Tail.SetClass(a.class + "-tail")
		a.Tail.WriteSVG(out)
		w.Print("</g>\n")
	}
	if a.Head != nil {
		w.Printf(`<g transform="rotate(%v %v %v)">`, a.angle()+90, x2, y2)
		alignHead(a.Head, x2, y2)
		a.Head.SetClass(a.class + "-head")
		a.Head.WriteSVG(out)
		w.Print("</g>\n")
	}
	return *err
}

func alignTail(s Shape, x, y int) {
	s.SetX(x)
	s.SetY(y - s.Height()/2)
}

func alignHead(s Shape, x, y int) {
	s.SetX(x - s.Width()/2) // specific to triangle
	s.SetY(y)
}

// AbsAngle
func (a *Line) AbsAngle() int { return int(a.absAngle()) }

func (a *Line) absAngle() float64 {
	return math.Abs(float64(a.angle()))
}

// Angle returns value in degrees. Right = 0, down = 90, left: 180, up = -90
func (a *Line) Angle() int { return a.angle() }

// angle returns degrees the head of an arrow should rotate depending
// on direction
func (a *Line) angle() int {
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
func (a *Line) DirQ1() bool {
	start, end := a.endpoints()
	return start.LeftOf(end) && end.Below(start)
}

// DirQ2 returns true if the arrow points to the bottom-left
// quadrant.
func (a *Line) DirQ2() bool {
	start, end := a.endpoints()
	return start.RightOf(end) && end.Below(start)
}

// DirQ3 returns true if the arrow points to the top-left
// quadrant.
func (a *Line) DirQ3() bool {
	start, end := a.endpoints()
	return start.RightOf(end) && end.Above(start)
}

// DirQ4 returns true if the arrow points to the top-right
// quadrant.
func (a *Line) DirQ4() bool {
	start, end := a.endpoints()
	return start.LeftOf(end) && end.Above(start)
}

func (a *Line) endpoints() (xy.Point, xy.Point) {
	return a.Start, a.End
}

func radians2degrees(A float64) int {
	return int(A * 180 / math.Pi)
}

func (a *Line) Height() int {
	v := []int{intAbs(a.Start.Y - a.End.Y)}
	if a.Head != nil {
		v = append(v, a.Head.Height())
	}
	if a.Tail != nil {
		v = append(v, a.Tail.Height())
	}
	return maxOf(v...)
}

func (a *Line) Width() int {
	v := []int{intAbs(a.Start.X - a.End.X)}
	if a.Head != nil {
		v = append(v, a.Head.Width())
	}
	if a.Tail != nil {
		v = append(v, a.Tail.Width())
	}
	return maxOf(v...)
}

func maxOf(v ...int) int {
	var max int
	for _, v := range v {
		if v < max {
			continue
		}
		max = v
	}
	return max
}

// ----------------------------------------

func (a *Line) Position() (int, int) {
	x, y := a.Start.XY()
	if a.End.LeftOf(a.Start) {
		x = a.End.X
	}
	if a.End.Above(a.Start) {
		y = a.End.Y
	}

	return x, y
}

// CenterPosition returns the center x, y values
func (a *Line) CenterPosition() (int, int) {
	x, y := a.Position()
	return x + a.Width()/2, y + a.Height()/2
}

func (a *Line) SetX(x int) {
	diff := a.Start.X - x
	a.Start.X = x
	a.End.X = a.End.X - diff // Set X2 so the entire arrow moves
}

func (a *Line) SetY(y int) {
	diff := a.Start.Y - y
	a.Start.Y = y
	a.End.Y = a.End.Y - diff // Set Y2 so the entire arrow moves
}

// Direction returns vertical or horizontal direction, Other if at an angle.
// If Other, use arrow.DirQn() methods to check to which quadrant.
func (a *Line) Direction() Direction {
	return NewDirection(a.Start, a.End)
}

func (a *Line) SetClass(c string) { a.class = c }

func NewArrowBetween(a, b Shape) *Line {
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
