package shape

import "io"

type Shape interface {
	// Position returns the xy position of the top left corner.
	Position() (x int, y int)
	SetX(int)
	SetY(int)
	Width() int
	Height() int
	// Direction returns in which direction the shape is drawn.
	// The direction and position is needed when aligning shapes.
	Direction() Direction
	SetClass(string)
	WriteSvg(io.Writer) error
}
