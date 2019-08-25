package shape

type Aligner int

func (Aligner) HAlignCenter(shapes ...Shape) { hAlign(Center, shapes...) }
func (Aligner) HAlignTop(shapes ...Shape)    { hAlign(Top, shapes...) }
func (Aligner) HAlignBottom(shapes ...Shape) { hAlign(Bottom, shapes...) }

func hAlign(adjust Adjust, objects ...Shape) {
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

func (Aligner) VAlignCenter(shapes ...Shape) { vAlign(Center, shapes...) }
func (Aligner) VAlignLeft(shapes ...Shape)   { vAlign(Left, shapes...) }
func (Aligner) VAlignRight(shapes ...Shape)  { vAlign(Right, shapes...) }

func vAlign(adjust Adjust, objects ...Shape) {
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

type Adjust int

const (
	Top Adjust = iota
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
