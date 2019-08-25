package shape

func NewAdjuster(s Shape) *Adjuster {
	return &Adjuster{
		shape:        s,
		defaultSpace: 30,
	}
}

type Adjuster struct {
	shape        Shape
	defaultSpace int
}

func (adjust *Adjuster) At(x, y int) {
	adjust.shape.SetX(x)
	adjust.shape.SetY(y)
}

func (adjust *Adjuster) RightOf(o Shape, space ...int) {
	x, _ := o.Position()
	adjust.shape.SetX(x + o.Width() + adjust.space(space))
}

func (adjust *Adjuster) Below(o Shape, space ...int) {
	x, y := o.Position()
	adjust.shape.SetY(y + o.Height() + adjust.space(space))
	adjust.shape.SetX(x)
}

func (adjust *Adjuster) space(space []int) int {
	if len(space) == 0 {
		return adjust.defaultSpace
	}
	return space[0]
}
