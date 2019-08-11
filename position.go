package design

type Positioned interface {
	SetX(x int)
	SetY(y int)
}

type PositionedDrawable interface {
	Positioned
	Drawable
}
