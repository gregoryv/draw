package shape

func AlignHorizontal(adjust Adjust, objects ...Shape) {
	mustAlign(adjust, objects, Top, Bottom, Center)
	first := objects[0]
	_, y := first.Position()
	for _, shape := range objects[1:] {
		switch adjust {
		case Top:
			shape.SetY(y)
		case Bottom:
			shape.SetY(y + first.Height() - shape.Height())
		case Center:
			shape.SetY(y + (first.Height()-shape.Height())/2)
		}
	}
}

func AlignVertical(adjust Adjust, objects ...Shape) {
	mustAlign(adjust, objects, Left, Right, Center)
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

func mustAlign(adjust Adjust, objects []Shape, ok ...Adjust) {
	if len(objects) < 2 {
		panic("Align must have 2 or more objects as arguments")
	}
	for _, a := range ok {
		if adjust == a {
			return
		}
	}
	panic("Cannot adjust Left or Right when horizontal.")
}

type Adjust int

const (
	Top Adjust = iota
	Left
	Right
	Bottom
	Center
)

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
	_, y := o.Position()
	adjust.shape.SetY(y + o.Height() + adjust.Space(l))
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
