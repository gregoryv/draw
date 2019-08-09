package design

type Element int

// Elements have been parsed from https://www.w3.org/TR/SVG11/eltindex.html
const (
	Element_undefined Element = iota // Undefined

	Element_rect // rect

	Element_last // Undefined
)

var elementNames map[Element]string = map[Element]string{
	Element_rect: "rect",
}

func (i Element) String() string {
	name, found := elementNames[i]
	if !found {
		return "undefined"
	}
	return name
}
