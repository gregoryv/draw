package shape

// Aligner type aligns multiple shapes
type Aligner int

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
			firstHigher := first.Height() > shape.Height()
			diff := intAbs(first.Height()-shape.Height()) / 2
			if shape, ok := shape.(*Label); ok {
				// labels are drawn from bottom left corner
				if firstHigher {
					diff += shape.Height()
				} else {
					diff -= shape.Height()
				}
			}
			switch {
			case firstHigher:
				shape.SetY(y + diff)
			case !firstHigher:
				shape.SetY(y - diff)
			}
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
			if first.Direction() == RL {
				shape.SetX(x - (first.Width()+shape.Width())/2)
			} else {
				shape.SetX(x + (first.Width()-shape.Width())/2)
			}

		}
	}
}

type Alignment int

const (
	Top Alignment = iota
	Left
	Right
	Bottom
	Center
)

type Direction int

const (
	Horizontal Direction = iota
	Vertical
	LR
	RL
)
