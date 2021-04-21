package shape

import "github.com/gregoryv/draw/xy"

// Aligner type aligns multiple shapes
type Aligner struct{}

// HAlignCenter aligns shape[1:] to shape[0] center coordinates horizontally
func (Aligner) HAlignCenter(shapes ...Shape) { hAlign(Center, shapes...) }

// HAlignTop aligns shape[1:] to shape[0] top coordinates horizontally
func (Aligner) HAlignTop(shapes ...Shape) { hAlign(Top, shapes...) }

// HAlignBottom aligns shape[1:] to shape[0] bottom coordinates horizontally
func (Aligner) HAlignBottom(shapes ...Shape) { hAlign(Bottom, shapes...) }

func hAlign(adjust Alignment, objects ...Shape) {
	first := objects[0]
	_, y := first.Position()
	for _, shape := range objects[1:] {
		switch adjust {
		case Top:
			shape.SetY(y)
		case Bottom:
			shape.SetY(y + first.Height() - shape.Height())
		case Center:
			diff := (first.Height() - shape.Height()) / 2
			shape.SetY(y + diff)
		}
	}
}

// VAlignCenter aligns shape[1:] to shape[0] center coordinates vertically
func (Aligner) VAlignCenter(shapes ...Shape) { vAlign(Center, shapes...) }

// VAlignLeft aligns shape[1:] to shape[0] left coordinates vertically
func (Aligner) VAlignLeft(shapes ...Shape) { vAlign(Left, shapes...) }

// VAlignRight aligns shape[1:] to shape[0] right coordinates vertically
func (Aligner) VAlignRight(shapes ...Shape) { vAlign(Right, shapes...) }

func vAlign(adjust Alignment, objects ...Shape) {
	first := objects[0]
	x, _ := first.Position()
	for _, shape := range objects[1:] {
		switch adjust {
		case Left:
			shape.SetX(x)
		case Right:
			shape.SetX(x + first.Width() - shape.Width())
		case Center:
			shape.SetX(x + (first.Width()-shape.Width())/2)
		}
	}
}

func centerAt(s Shape, p xy.Point) {
	x, y := p.XY()
	centerXY(s, x, y)
}

func centerXY(s Shape, x, y int) {
	s.SetX(x - s.Width()/2)
	s.SetY(y - s.Height()/2)
}

type Alignment int

const (
	Top Alignment = iota
	Left
	Right
	Bottom
	Center
)

func NewDirection(from, to xy.Point) Direction {
	switch {
	case from.LeftOf(to) && from.Y == to.Y:
		return DirectionRight
	case from.LeftOf(to) && from.Above(to):
		return DirectionDownRight
	case from.Above(to) && from.X == to.X:
		return DirectionDown
	case from.RightOf(to) && from.Above(to):
		return DirectionDownLeft
	case from.RightOf(to) && from.Y == to.Y:
		return DirectionLeft
	case from.Below(to) && from.RightOf(to):
		return DirectionUpLeft
	case from.Below(to) && from.X == to.X:
		return DirectionUp
	default: // from.LeftOf(to) && from.Below(to):
		return DirectionUpRight
	}
}

type Direction uint

const (
	DirectionRight Direction = (1 << iota)
	DirectionLeft
	DirectionUp
	DirectionDown

	DirectionDownRight = DirectionDown | DirectionRight
	DirectionDownLeft  = DirectionDown | DirectionLeft
	DirectionUpLeft    = DirectionUp | DirectionLeft
	DirectionUpRight   = DirectionUp | DirectionRight
)

// Method
func (d Direction) Is(dir Direction) bool {
	return (d & dir) == dir
}
