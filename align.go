package design

func AlignHorizontal(adjust Adjust, objects ...Positioned) {
	mustAlign(adjust, objects, Top, Bottom, Center)
	first := objects[0]
	y := first.Y()
	for _, obj := range objects[1:] {
		switch adjust {
		case Top:
			obj.SetY(y)
		case Bottom:
			obj.SetY(y + first.Height() - obj.Height())
		case Center:
			obj.SetY(y + (first.Height()-obj.Height())/2)
		}
	}
}

func AlignVertical(adjust Adjust, objects ...Positioned) {
	mustAlign(adjust, objects, Left, Right, Center)
	first := objects[0]
	x := first.X()
	for _, obj := range objects[1:] {
		switch adjust {
		case Left:
			obj.SetX(x)
		case Right:
			obj.SetX(x + first.Width() - obj.Width())
		case Center:
			obj.SetX(x + (first.Width()-obj.Width())/2)
		}
	}
}

func mustAlign(adjust Adjust, objects []Positioned, ok ...Adjust) {
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
	obj          Positioned
	defaultSpace int
}

func (adjust *Adjuster) At(x, y int) {
	adjust.obj.SetX(x)
	adjust.obj.SetY(y)
}

func (adjust *Adjuster) RightOf(o Positioned, l ...int) {
	adjust.obj.SetX(o.X() + o.Width() + adjust.Space(l))
}

func (adjust *Adjuster) Below(o Positioned, l ...int) {
	adjust.obj.SetY(o.Y() + o.Height() + adjust.Space(l))
}

func (adjust *Adjuster) Space(space []int) int {
	if len(space) == 0 {
		return adjust.defaultSpace
	}
	return space[0]
}
