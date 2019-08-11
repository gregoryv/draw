package svg

import (
	"strconv"

	"github.com/gregoryv/go-design/xml"
)

type Element int

// Elements have been parsed from https://www.w3.org/TR/SVG11/eltindex.html
const (
	Element_undefined Element = iota // Undefined

	Element_line
	Element_rect
	Element_text

	Element_last // Undefined
)

var elementNames map[Element]string = map[Element]string{
	Element_line: "line",
	Element_rect: "rect",
	Element_text: "text",
}

func (i Element) String() string {
	name, found := elementNames[i]
	if !found {
		return "undefined"
	}
	return name
}

// Defined SVG attributes as constructor functions

func attr(key, val string) xml.Attribute {
	return xml.NewAttribute(key, val)
}

func xp(v int) xml.Attribute      { return attr("x", strconv.Itoa(v)) }
func yp(v int) xml.Attribute      { return attr("y", strconv.Itoa(v)) }
func width(v int) xml.Attribute   { return attr("width", strconv.Itoa(v)) }
func height(v int) xml.Attribute  { return attr("height", strconv.Itoa(v)) }
func fill(v string) xml.Attribute { return attr("fill", v) }

func Rect(x, y, w, h int, attr ...xml.Attribute) *xml.Node {
	all := append([]xml.Attribute{xp(x), yp(y), width(w), height(h)},
		attr...)
	return xml.NewNode(Element_rect, all...)
}

func Text(x, y int, s string, attr ...xml.Attribute) *xml.Node {
	all := append([]xml.Attribute{xp(x), yp(y)}, attr...)
	text := xml.NewNode(Element_text, all...)
	text.Append(xml.CData(s))
	return text
}
