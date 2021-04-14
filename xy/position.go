// Package xy provides xy Position
package xy

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) LeftOf(q Point) bool  { return p.X < q.X }
func (p Point) RightOf(q Point) bool { return p.X > q.X }
func (p Point) Above(q Point) bool   { return p.Y < q.Y }
func (p Point) Below(q Point) bool   { return p.Y > q.Y }
func (p Point) XY() (int, int)       { return p.X, p.Y }

func (p Point) String() string {
	x, y := p.XY()
	return fmt.Sprintf("%v,%v", x, y)
}

func (p Point) XYfloat64() (float64, float64) {
	return float64(p.X), float64(p.Y)
}

func (p Point) Distance(q Point) float64 {
	x1, y1 := p.XYfloat64()
	x2, y2 := q.XYfloat64()
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
