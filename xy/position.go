package xy

type Position struct {
	X, Y int
}

func (p *Position) LeftOf(q Position) bool {
	return p.X < q.X
}

func (p *Position) RightOf(q Position) bool {
	return p.X > q.X
}

func (p *Position) Above(q Position) bool {
	return p.Y < q.Y
}

func (p *Position) Below(q Position) bool {
	return p.Y > q.Y
}

func (p *Position) XY() (int, int) { return p.X, p.Y }
