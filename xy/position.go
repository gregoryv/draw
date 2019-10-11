// Package xy provides xy Position
package xy

import (
	"fmt"
	"math"
)

type Position struct {
	X, Y int
}

func (p Position) Equals(q Position) bool {
	return p.X == q.X && p.Y == q.Y
}

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

func (p Position) Distance(q Position) float64 {
	x1, y1 := p.XYfloat64()
	x2, y2 := q.XYfloat64()
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
