package shape

import "io"

type Shape interface {
	Box
	SetX(int)
	SetY(int)

	// Direction returns in which direction the shape is drawn.
	// The direction and position is needed when aligning shapes.
	Direction() Direction
	SetClass(string)
	WriteSVG(io.Writer) error
}

func SetClass(class string, shapes ...Shape) {
	for _, s := range shapes {
		s.SetClass(class)
	}
}
