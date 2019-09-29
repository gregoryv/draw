package shape

type Shape interface {
	// Position returns the starting point of a shape.
	// Mostly the top left corner.
	Position() (x int, y int)
	SetX(int)
	SetY(int)
	Width() int
	Height() int

	// Direction returns in which direction the shape is drawn.
	// The direction and position is needed when aligning shapes.
	Direction() Direction
}
