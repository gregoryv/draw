package design

import (
	"github.com/gregoryv/go-design/xml"
)

type Element int

// Elements have been parsed from https://www.w3.org/TR/SVG11/eltindex.html
const (
	Element_undefined Element = iota // Undefined

	Element_rect // rect
	Element_text // text

	Element_last // Undefined
)

var elementNames map[Element]string = map[Element]string{
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

func x(v string) xml.Attribute      { return attr("x", v) }
func y(v string) xml.Attribute      { return attr("y", v) }
func width(v string) xml.Attribute  { return attr("width", v) }
func height(v string) xml.Attribute { return attr("height", v) }
func style(v string) xml.Attribute  { return attr("style", v) }
func fill(v string) xml.Attribute   { return attr("fill", v) }
