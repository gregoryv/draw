package shape

// NewAdjuster returns an adjuster with default space of 30 pixels.
func NewAdjuster(s Shape) *Adjuster {
	return &Adjuster{
		shape:        s,
		defaultSpace: 30,
	}
}

// Adjuster is used to position a shape relative to other shapes or at
// a specific xy position.
type Adjuster struct {
	shape        Shape
	defaultSpace int
}

// At sets the x, y coordinates of the wrapped shape
func (adjust *Adjuster) At(x, y int) {
	adjust.shape.SetX(x)
	adjust.shape.SetY(y)
}

// RightOf places the wrapped shape to the right of o. Optional space
// to override default.
func (adjust *Adjuster) RightOf(o Shape, space ...int) {
	x, _ := o.Position()
	adjust.shape.SetX(x + o.Width() + adjust.space(space))
}

// Below places the wrapped shape below o. Optional space to override
// default.
func (adjust *Adjuster) Below(o Shape, space ...int) {
	x, y := o.Position()
	adjust.shape.SetY(y + o.Height() + adjust.space(space))
	adjust.shape.SetX(x)
}

// Above places the wrapped shape above o. Optional space to override
// default.
func (adjust *Adjuster) Above(o Shape, space ...int) {
	x, y := o.Position()
	adjust.shape.SetY(y - (o.Height() + adjust.space(space)))
	adjust.shape.SetX(x)
}

func (adjust *Adjuster) space(space []int) int {
	if len(space) == 0 {
		return adjust.defaultSpace
	}
	return space[0]
}
