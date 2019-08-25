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

func NewAdjuster(s Shape) *Adjuster {
	return &Adjuster{
		shape:        s,
		defaultSpace: 30,
	}
}

type Adjuster struct {
	shape        Shape
	defaultSpace int
}

func (adjust *Adjuster) At(x, y int) {
	adjust.shape.SetX(x)
	adjust.shape.SetY(y)
}

func (adjust *Adjuster) RightOf(o Shape, l ...int) {
	x, _ := o.Position()
	adjust.shape.SetX(x + o.Width() + adjust.Space(l))
}

func (adjust *Adjuster) Below(o Shape, l ...int) {
	x, y := o.Position()
	adjust.shape.SetY(y + o.Height() + adjust.Space(l))
	adjust.shape.SetX(x)
}

func (adjust *Adjuster) Space(space []int) int {
	if len(space) == 0 {
		return adjust.defaultSpace
	}
	return space[0]
}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
	LR
	RL
)

type Shape interface {
	// Position returns the center x,y values
	Position() (x int, y int)
	SetX(int)
	SetY(int)
	Width() int
	Height() int
	Direction() Direction
}
