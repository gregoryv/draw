package xy

import "fmt"

func NewLine(x1, y1, x2, y2 int) *Line {
	return &Line{
		Position{x1, y1},
		Position{x2, y2},
	}
}

type Line struct {
	Start, End Position
}

func (l1 *Line) String() string {
	return fmt.Sprint(l1.Start, " -- ", l1.End)
}

//https://en.wikipedia.org/wiki/Line-line_intersection

// todo this calculates elongated l1 and l2, I need segment
// intersections
func (l1 *Line) Intersect(l2 *Line) (Position, error) {
	x1, y1 := l1.Start.XYfloat64()
	x2, y2 := l1.End.XYfloat64()
	x3, y3 := l2.Start.XYfloat64()
	x4, y4 := l2.End.XYfloat64()

	tn := (x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)
	d := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	t := tn / d

	un := (x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)
	u := -1 * (un / d)

	P := Position{}
	fmt.Println(d, t, u)
	switch {
	case 0.0 <= t && t <= 1.0:
		P.X = int(x1 + t*(x2-x1))
		P.Y = int(y1 + t*(y2-y1))
	case 0.0 <= u && u <= 1.0:
		P.X = int(x3 + u*(x4-x3))
		P.Y = int(y3 + u*(y4-y3))
	default:
		return P, fmt.Errorf("Not intersecting")
	}

	return P, nil
}

func (l1 *Line) IntersectSegment(l2 *Line) (Position, error) {
	p, err := l1.Intersect(l2)
	if err != nil {
		return p, err
	}
	if !l1.Contains(p) || !l2.Contains(p) {
		return p, fmt.Errorf("Not intersecting")
	}
	return p, nil
}

func (l1 *Line) Contains(p Position) bool {
	if l1.Start.X != l1.End.X { // not vertical
		switch {
		case l1.Start.X <= p.X && p.X <= l1.End.X:
			return true
		case l1.Start.X >= p.X && p.X >= l1.End.X:
			return true
		}
	} else {
		switch {
		case l1.Start.Y <= p.X && p.Y <= l1.End.Y:
			return true
		case l1.Start.Y >= p.Y && p.Y >= l1.End.Y:
			return true
		}
	}
	return false
}
