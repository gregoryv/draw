package design

type Pos struct {
	x, y int
}

func (pos *Pos) X() int { return pos.x }
func (pos *Pos) Y() int { return pos.y }

func (pos *Pos) SetX(x int) { pos.x = x }
func (pos *Pos) SetY(y int) { pos.y = y }
