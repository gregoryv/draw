package shape

import "io"

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
	SetClass(string)
	WriteSvg(io.Writer) error
}

func (r *Record) Position() (int, int) { return r.X, r.Y }
func (r *Record) SetX(x int)           { r.X = x }
func (r *Record) SetY(y int)           { r.Y = y }
func (r *Record) Direction() Direction { return LR }
func (r *Record) SetClass(c string)    { r.class = c }

func (t *Triangle) Position() (int, int) { return t.Pos.XY() }
func (t *Triangle) SetX(x int)           { t.Pos.X = x }
func (t *Triangle) SetY(y int)           { t.Pos.Y = y }
func (t *Triangle) Width() int           { return 8 }
func (t *Triangle) Height() int          { return 4 }
func (t *Triangle) Direction() Direction { return LR }
func (t *Triangle) SetClass(c string)    { t.class = c }
