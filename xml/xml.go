package xml

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	element    fmt.Stringer
	attributes Attributes
	children   []Drawable
}

type Drawable interface {
	WriteTo(io.Writer) (int, error)
}

func NewNode(element fmt.Stringer, attributes ...Attribute) *Node {
	return &Node{
		element:    element,
		attributes: attributes,
		children:   make([]Drawable, 0),
	}
}

func (node *Node) Append(child *Node) {
	node.children = append(node.children, child)
}

func (node *Node) WriteTo(w io.Writer) (int, error) {
	var total int
	if !node.HasChildren() {
		return fmt.Fprintf(w, "<%s%s/>\n", node.element, node.attributes)
	}
	n, err := node.writeOpenTo(w)
	if err != nil {
		return n, err
	}
	total += n
	for _, child := range node.children {
		n, err := child.WriteTo(w)
		if err != nil {
			return total + n, err
		}
		total += n
	}
	n, err = node.writeCloseTo(w)
	return total + n, err
}

func (node *Node) writeOpenTo(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "<%s%s>", node.element, node.attributes)
}

func (node *Node) writeCloseTo(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "</%s>", node.element)
}

func (node *Node) HasChildren() bool {
	return len(node.children) > 0
}

func NewAttribute(key, val string) Attribute {
	return Attribute{key, val}
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
