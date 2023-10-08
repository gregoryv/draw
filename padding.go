package draw

type Padding struct {
	Left, Top, Right, Bottom int
}

func (p *Padding) SetScale(v float64) {
	p.Left = int(float64(p.Left) * v)
	p.Top = int(float64(p.Top) * v)
	p.Right = int(float64(p.Right) * v)
	p.Bottom = int(float64(p.Bottom) * v)
}
