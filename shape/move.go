package shape

func Move(m Movable, xd, yd int) {
	x, y := m.Position()
	m.SetX(x + xd)
	m.SetY(y + yd)
}

type Movable interface {
	Position() (x int, y int)
	SetX(int)
	SetY(int)
}
