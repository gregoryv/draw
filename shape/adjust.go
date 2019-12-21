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
func (a *Adjuster) At(x, y int) {
	a.shapes[0].SetX(x)
	a.shapes[0].SetY(y)
}

// RightOf places the wrapped shape to the right of o. Optional space
// to override default.
func (a *Adjuster) RightOf(o Shape, space ...int) {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetX(x + next.Width() + a.space(space))
		s.SetY(y)
		next = s
	}
}

// LeftOf places the wrapped shape to the left of o. Optional space
// to override default.
func (a *Adjuster) LeftOf(o Shape, space ...int) {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetX(x - (next.Width() + a.space(space)))
		s.SetY(y)
		next = s
	}
}

// Below places the wrapped shape below o. Optional space to override
// default.
func (a *Adjuster) Below(o Shape, space ...int) {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetY(y + next.Height() + a.space(space))
		s.SetX(x)
		next = s
	}
}

// Above places the wrapped shape above o. Optional space to override
// default.
func (a *Adjuster) Above(o Shape, space ...int) {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetY(y - (next.Height() + a.space(space)))
		s.SetX(x)
		next = s
	}
}

func (a *Adjuster) space(space []int) int {
	if len(space) == 0 {
		return a.defaultSpace
	}
	return space[0]
}
