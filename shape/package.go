package shape

type Stringer interface {
	String() string
}

type svg interface {
	Svg() string
}
