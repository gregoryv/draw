package shape

// NewAdjuster returns an adjuster with default space of 30 pixels.
func NewAdjuster(s ...Shape) *Adjuster {
	return &Adjuster{
		shapes:       s,
		defaultSpace: 30,
	}
}

// Adjuster is used to position a shape relative to other shapes or at
// a specific xy position.
type Adjuster struct {
	shapes       []Shape
	defaultSpace int
}

// At sets the x, y coordinates of the wrapped shape
func (adjust *Adjuster) At(x, y int) {
	adjust.shapes[0].SetX(x)
	adjust.shapes[0].SetY(y)
}

// RightOf places the wrapped shape to the right of o. Optional space
// to override default.
func (adjust *Adjuster) RightOf(o Shape, space ...int) {
	next := o
	for _, s := range adjust.shapes {
		x, y := next.Position()
		s.SetX(x + next.Width() + adjust.space(space))
		s.SetY(y)
		next = s
	}
}

// LeftOf places the wrapped shape to the left of o. Optional space
// to override default.
func (adjust *Adjuster) LeftOf(o Shape, space ...int) {
	next := o
	for _, s := range adjust.shapes {
		x, y := next.Position()
		s.SetX(x - (next.Width() + adjust.space(space)))
		s.SetY(y)
		next = s
	}
}

// Below places the wrapped shape below o. Optional space to override
// default.
func (adjust *Adjuster) Below(o Shape, space ...int) {
	next := o
	for _, s := range adjust.shapes {
		x, y := next.Position()
		s.SetY(y + next.Height() + adjust.space(space))
		s.SetX(x)
		next = s
	}
}

// Above places the wrapped shape above o. Optional space to override
// default.
func (adjust *Adjuster) Above(o Shape, space ...int) {
	next := o
	for _, s := range adjust.shapes {
		x, y := next.Position()
		s.SetY(y - (next.Height() + adjust.space(space)))
		s.SetX(x)
		next = s
	}
}

func (adjust *Adjuster) space(space []int) int {
	if len(space) == 0 {
		return adjust.defaultSpace
	}
	return space[0]
}
