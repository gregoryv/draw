package design

type Element int

//go:generate stringer -linecomment -type Element element.go

// Elements have been parsed from https://www.w3.org/TR/SVG11/eltindex.html
const (
	Element_undefined Element = iota // Undefined

	Element_rect // rect

	Element_last // Undefined
)
