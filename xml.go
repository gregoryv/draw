package design

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	element    Element
	attributes Attributes
}

func NewNode(element Element, attributes ...Attribute) *Node {
	return &Node{
		element:    element,
		attributes: attributes,
	}
}

func (node *Node) WriteTo(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "<%s%s>", node.element, node.attributes)
}

type Attribute struct {
	key, val string
}

func (attr *Attribute) String() string {
	return fmt.Sprintf("%s=%q", attr.key, attr.val)
}

type Attributes []Attribute

func (attributes Attributes) String() string {
	if len(attributes) == 0 {
		return ""
	}
	all := make([]string, len(attributes))
	for i, attr := range attributes {
		all[i] = attr.String()
	}
	return " " + strings.Join(all, " ")
}
