package shape

type Shape interface {
	Position() (x int, y int)
	SetX(int)
	SetY(int)
	Width() int
	Height() int
	Direction() Direction
}
