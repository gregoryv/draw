package shape

// NewAdjuster returns an adjuster using DefaultSpacing.
func NewAdjuster(s ...Shape) *Adjuster {
	return &Adjuster{
		shapes:  s,
		Spacing: DefaultSpacing,
	}
}

// Adjuster is used to position a shape relative to other shapes or at
// a specific xy position.
type Adjuster struct {
	shapes  []Shape
	Spacing int
}

// At sets the x, y coordinates of the wrapped shape
func (a *Adjuster) At(x, y int) *Adjuster {
	a.shapes[0].SetX(x)
	a.shapes[0].SetY(y)
	return a
}

// Move adjusts shapes by moving them +/- in x and or y direction
func (a *Adjuster) Move(dx, dy int) *Adjuster {
	for i, _ := range a.shapes {
		Move(a.shapes[i], dx, dy)
	}
	return a
}

// RightOf places the wrapped shape to the right of o. Optional space
// to override default.
func (a *Adjuster) RightOf(o Shape, space ...int) *Adjuster {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetX(x + next.Width() + a.space(space))
		s.SetY(y)
		next = s
	}
	return a
}

// LeftOf places the wrapped shape to the left of o. Optional space
// to override default.
func (a *Adjuster) LeftOf(o Shape, space ...int) *Adjuster {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetX(x - (next.Width() + a.space(space)))
		s.SetY(y)
		next = s
	}
	return a
}

// Below places the wrapped shape below o. Optional space to override
// default.
func (a *Adjuster) Below(o Shape, space ...int) *Adjuster {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetY(y + next.Height() + a.space(space))
		s.SetX(x)
		next = s
	}
	return a
}

// Above places the wrapped shape above o. Optional space to override
// default.
func (a *Adjuster) Above(o Shape, space ...int) *Adjuster {
	next := o
	for _, s := range a.shapes {
		x, y := next.Position()
		s.SetY(y - (next.Height() + a.space(space)))
		s.SetX(x)
		next = s
	}
	return a
}

func (a *Adjuster) space(space []int) int {
	if len(space) == 0 {
		return a.Spacing
	}
	return space[0]
}
