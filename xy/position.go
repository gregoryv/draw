// Package xy provides xy Position
package xy

import "fmt"

type Position struct {
	X, Y int
}

func (p Position) Equals(q Position) bool  { return p.X == q.X && p.Y == q.Y }
func (p Position) LeftOf(q Position) bool  { return p.X < q.X }
func (p Position) RightOf(q Position) bool { return p.X > q.X }
func (p Position) Above(q Position) bool   { return p.Y < q.Y }
func (p Position) Below(q Position) bool   { return p.Y > q.Y }
func (p Position) XY() (int, int)          { return p.X, p.Y }

func (p Position) String() string {
	x, y := p.XY()
	return fmt.Sprintf("%v,%v", x, y)
}

func (p Position) XYfloat64() (float64, float64) {
	return float64(p.X), float64(p.Y)
}

func (p Position) Within(r Rect) bool {
	if r.TopLeft.LeftOf(r.BottomRight) {
		withinXspan := (r.TopLeft.X <= p.X && p.X <= r.BottomRight.X)
		withinYspan := (r.TopLeft.Y <= p.Y && p.Y <= r.BottomRight.Y)
		return withinXspan && withinYspan
	}
	withinXspan := (r.TopLeft.X >= p.X && p.X >= r.BottomRight.X)
	withinYspan := (r.TopLeft.Y >= p.Y && p.Y >= r.BottomRight.Y)
	return withinXspan && withinYspan
}
