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
